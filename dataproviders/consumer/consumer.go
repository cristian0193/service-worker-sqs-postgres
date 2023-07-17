package consumer

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
	domain2 "service-worker-sqs-postgres/core/domain"
	"service-worker-sqs-postgres/core/domain/entity"
	"service-worker-sqs-postgres/dataproviders/awssqs"
	repository "service-worker-sqs-postgres/dataproviders/repository/events"
	"sync"
	"time"
)

// SQSSource event stream representation to SQS.
type SQSSource struct {
	sqs         *awssqs.ClientSQS
	log         *zap.SugaredLogger
	maxMessages int
	closed      bool
	repo        repository.IEventsRepository
	wg          sync.WaitGroup
}

// New return an event stream instance from SQS.
func New(sqsClient *awssqs.ClientSQS, logger *zap.SugaredLogger, maxMessages int, repo repository.IEventsRepository) (*SQSSource, error) {
	return &SQSSource{
		sqs:         sqsClient,
		log:         logger,
		maxMessages: maxMessages,
		repo:        repo,
		wg:          sync.WaitGroup{},
	}, nil
}

// Consume opens a channel and sends entities created from SQS messages.
func (s *SQSSource) Consume() <-chan *domain2.Event {
	out := make(chan *domain2.Event, s.maxMessages)
	go func() {
		for {
			if s.closed {
				break
			}
			messages, err := s.sqs.GetMessages()
			if err != nil {
				s.log.Errorf("Error getting messages from SQS: %v", err)
				continue
			}
			if len(messages) == 0 {
				s.log.Debug("No messages found from SQS")
			}
			for _, msg := range messages {
				s.processMessage(msg, out)
			}
			s.wg.Wait()
		}
		close(out)
	}()

	return out
}

// processMessage read message in queue.
func (s *SQSSource) processMessage(msg *sqs.Message, out chan *domain2.Event) {
	var records domain2.Events
	err := json.Unmarshal([]byte(*msg.Body), &records)
	if err != nil {
		s.log.Errorf("Error processing message from SQS: %v", err)
		return
	}
	retry := "0"
	val, ok := msg.Attributes[sqs.MessageSystemAttributeNameApproximateReceiveCount]
	if ok {
		retry = *val
	}

	logger := s.log.With("retry", retry)
	logger.Infof("Step 1 - Start to process SQS event")

	eventDB := &entity.Events{
		ID:      *msg.MessageId,
		Message: records.Message,
		Date:    time.Now().Format(time.RFC3339),
	}

	if err = s.repo.Insert(eventDB); err != nil {
		logger.Errorf("Error inserting message: %v", err)
	}
	logger.Info("Step 2 - Event saved in postgres")

	event := &domain2.Event{
		ID:            *msg.MessageId,
		Retry:         retry,
		Records:       records,
		OriginalEvent: msg,
		Log:           s.log,
	}
	s.wg.Add(1)
	logger.Infof("Step 3 - Event produced for ID = %s)", event.ID)
	out <- event
}

// Processed notify that event of consolidate file was processed.
func (s *SQSSource) Processed(event *domain2.Event) error {
	defer s.wg.Done()
	logger := event.Log

	if events, ok := event.OriginalEvent.(*sqs.Message); ok {
		if err := s.sqs.DeleteMessage(events); err != nil {
			logger.Errorf("error deleting of sqs message. %v", err)
			return err
		}
		logger.Infof("Step 4 - Successful deleted sqs message")
		return nil
	}
	logger.Warnf("Event isn't sqs message")
	return nil
}

// Close the event stream.
func (s *SQSSource) Close() error {
	s.closed = true
	s.wg.Wait()
	return nil
}
