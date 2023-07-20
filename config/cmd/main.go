package main

import (
	"os"
	"os/signal"
	"service-worker-sqs-postgres/config/cmd/builder"
	cases "service-worker-sqs-postgres/core/usecases/events"
	repository "service-worker-sqs-postgres/dataproviders/postgres/repository/events"
	"service-worker-sqs-postgres/dataproviders/server"
	"service-worker-sqs-postgres/entrypoints/controllers/events"
	"syscall"
)

func main() {

	// logger is initialized
	logger := builder.NewLogger()
	logger.Info("Starting service-worker-sqs-postgres ...")
	defer builder.Sync(logger)

	// config is initialized
	config, err := builder.LoadConfig()
	if err != nil {
		logger.Fatalf("error in LoadConfig : %v", err)
	}

	// session aws is initialized
	session, err := builder.NewSession(config)
	if err != nil {
		logger.Fatalf("error in Session : %v", err)
	}

	// db is initialized
	db, err := builder.NewDB(config)
	if err != nil {
		logger.Fatalf("error in RDS : %v", err)
	}

	// repositories are initialized
	eventRepository := repository.NewEventRepository(db)

	// usecases are initialized
	eventUseCases := cases.NewEventUseCases(eventRepository)

	// controllers are initialized
	eventController := events.NewEventController(eventUseCases)

	// sqs is initialized
	sqs, err := builder.NewSQS(logger, config, session, eventRepository)
	if err != nil {
		logger.Fatalf("error in SQS : %v", err)
	}

	// processor is initialized
	processor, err := builder.NewProcessor(logger, sqs)
	if err != nil {
		logger.Fatalf("error in Processor : %v", err)
	}
	go processor.Start()

	// server is initialized
	srv := server.NewServer(config.Port, eventController)
	if err = srv.Start(); err != nil {
		logger.Fatalf("error Starting Server: %v", err)
	}

	// Graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	sig := <-sigQuit

	logger.Infof("Shutting down server with signal [%s] ...", sig.String())
	if err = sqs.Close(); err != nil {
		logger.Error("error Closing Consumer SQS: %v", err)
	}

	if err = srv.Stop(); err != nil {
		logger.Error("error Stopping Server: %v", err)
	}

	logger.Info("service-worker-sqs-postgres ended")

}
