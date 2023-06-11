package consumer

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
	"service-template-golang/clients/awssqs"
	"service-template-golang/domain"
	"sync"
)

// SQSSource event stream representation to SQS.
type SQSSource struct {
	sqs         *awssqs.ClientSQS
	log         *zap.SugaredLogger
	maxMessages int
	closed      bool
	wg          sync.WaitGroup
}

// New return a event stream instance from SQS.
func New(sqsClient *awssqs.ClientSQS, logger *zap.SugaredLogger, maxMessages int) (*SQSSource, error) {
	return &SQSSource{
		sqs:         sqsClient,
		log:         logger,
		maxMessages: maxMessages,
		wg:          sync.WaitGroup{},
	}, nil
}

// Consume opens a channel and sends entities created from SQS messages and their s3 files.
func (s *SQSSource) Consume() <-chan *domain.Event {
	out := make(chan *domain.Event, s.maxMessages)
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

// EventProcessed notify that event of consolidate file was processed.
func (s *SQSSource) EventProcessed() {
	s.wg.Done()
}

// Close the event stream.
func (s *SQSSource) Close() error {
	s.closed = true
	s.wg.Wait()
	return nil
}

func (s *SQSSource) processMessage(msg *sqs.Message, out chan *domain.Event) {
	var records []map[string]interface{}
	err := json.Unmarshal([]byte(*msg.Body), &records)
	if err != nil {
		s.log.Errorf("Error processing message from SQS: %v", err)
		return
	}
	// get retry number from message
	retry := "0"
	val, ok := msg.Attributes[sqs.MessageSystemAttributeNameApproximateReceiveCount]
	if ok {
		retry = *val
	}

	logger := s.log.With("retry", retry)
	logger.Infof("Start to process SQS event")

	event := &domain.Event{
		ID:            *msg.MessageId,
		Retry:         retry,
		Records:       records,
		OriginalEvent: msg,
		Log:           s.log,
	}
	s.wg.Add(1)
	logger.Infof("Event produced for ID = %s)", event.ID)
	out <- event
}
