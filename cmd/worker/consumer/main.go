package main

import (
	"context"
	"log"
	"sync"

	appConfig "github.com/FIAP-SOAT-G20/tc4-order-service/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/infrastructure/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

var (
	wg sync.WaitGroup
)

func main() {

	// Load AWS Config
	appCfg := appConfig.LoadConfig()

	loggerInstance := logger.NewLogger(appCfg.Environment)


	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			appCfg.AWS_SQSKey,
			appCfg.AWS_SQSSecret,
			appCfg.AWS_SQSToken)),
		config.WithRegion(appCfg.AWS_SQSRegion),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Using the Config value, create the SQS client
	sqsClient := sqs.NewFromConfig(cfg)

	loggerInstance.Info("SQS client created", "region", appCfg.AWS_SQSRegion, "url", appCfg.AWS_SQSURL)

	// Receive messages from SQS
	for {
		results, err := sqsClient.ReceiveMessage(
			context.Background(),
			&sqs.ReceiveMessageInput{
				QueueUrl:            &appCfg.AWS_SQSURL,
				MaxNumberOfMessages: 10,
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		for _, message := range results.Messages {

			wg.Add(1)

			go processJob(message, sqsClient, appCfg.AWS_SQSURL, loggerInstance)

			wg.Wait()
		}
	}
}

func processJob(message types.Message, sqsClient *sqs.Client, awsSQSURL string, logger *logger.Logger) {
	defer wg.Done()

	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(awsSQSURL),
		ReceiptHandle: message.ReceiptHandle,
	}
	// Process message here
	logger.Info("Processing message", "messageID", *message.MessageId, "body", *message.Body)

	_, err := sqsClient.DeleteMessage(context.TODO(), deleteParams)

	if err != nil {
		log.Fatal(err)
	}
}
