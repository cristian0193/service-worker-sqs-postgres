package main

import (
	"fmt"
	"service-template-golang/builder"
)

func main() {

	logger := builder.NewLogger()
	logger.Info("Starting service-template-golang ...")
	defer builder.Sync(logger)

	_, err := builder.LoadConfig()
	if err != nil {
		logger.Fatalf("Error in LoadConfig : %v", err)
	}

	consumer, err := builder.NewSQSConsumer(logger)
	if err != nil {
		logger.Fatalf("Error in Processor : %v", err)
	}

	processor, err := builder.NewProcessor(logger, consumer)
	if err != nil {
		logger.Fatalf("Error in Processor : %v", err)
	}
	go processor.Start()
	_ = processor.Stop()

	fmt.Print(processor)

}
