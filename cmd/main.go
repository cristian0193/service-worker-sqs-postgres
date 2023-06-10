package main

import (
	"os"
	"os/signal"
	"service-template-golang/builder"
	"syscall"
)

func main() {

	logger := builder.NewLogger()
	logger.Info("Starting service-template-golang ...")
	defer builder.Sync(logger)

	config, err := builder.LoadConfig()
	if err != nil {
		logger.Fatalf("Error in LoadConfig : %v", err)
	}

	consumer, err := builder.NewSQSConsumer(logger, config)
	if err != nil {
		logger.Fatalf("Error in Processor : %v", err)
	}

	processor, err := builder.NewProcessor(logger, consumer)
	if err != nil {
		logger.Fatalf("Error in Processor : %v", err)
	}
	go processor.Start()

	// Graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	sig := <-sigQuit

	logger.Infof("Shutting down server with signal [%s] ...", sig.String())
	if err = consumer.Close(); err != nil {
		logger.Error("consumer.Stop: %v", err)
	}

	logger.Info("service-template-golang ended")

}
