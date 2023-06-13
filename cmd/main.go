package main

import (
	"os"
	"os/signal"
	"service-template-golang/cmd/builder"
	"service-template-golang/http"
	"service-template-golang/http/controllers"
	"service-template-golang/http/repository"
	"service-template-golang/http/services"
	"syscall"
)

func main() {

	// logger is initialized
	logger := builder.NewLogger()
	logger.Info("Starting service-template-golang ...")
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

	// sqs is initialized
	sqs, err := builder.NewSQS(logger, config, session, db)
	if err != nil {
		logger.Fatalf("error in SQS : %v", err)
	}

	// processor is initialized
	processor, err := builder.NewProcessor(logger, sqs, db)
	if err != nil {
		logger.Fatalf("error in Processor : %v", err)
	}
	go processor.Start()

	// repositories are initialized
	eventsRepository := repository.NewEventsRepository(db)

	// services are initialized
	eventsService := services.NewEventsService(eventsRepository)

	// controllers are initialized
	eventsController := controllers.NewEventsController(eventsService)

	// server is initialized
	server := http.NewServer(config.Port, eventsController)
	if err = server.Start(); err != nil {
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

	if err = server.Stop(); err != nil {
		logger.Error("error Stopping Server: %v", err)
	}

	logger.Info("service-template-golang ended")

}
