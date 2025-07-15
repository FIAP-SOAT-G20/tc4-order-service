package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/core/domain"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/core/domain/entity"
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

	err := processedMessage(message, logger)
	if err != nil {
		logger.Error("failed to process message", "error", err.Error(), "messageID", *message.MessageId)
		// send the message to a dead-letter queue for further investigation
	}

	_, err = sqsClient.DeleteMessage(context.TODO(), deleteParams)

	if err != nil {
		log.Fatal(err)
	}
}

func processedMessage(message types.Message, logger *logger.Logger) error {
	// Here you can implement the logic to process the message
	// For example, you can unmarshal the message body and update the order status in your database
	logger.Info("Processing message", "messageID", *message.MessageId, "body", *message.Body)

	// Unmarshal the message body to your entity
	var updatedOrderStatus entity.OrderStatusUpdated
	err := json.Unmarshal([]byte(*message.Body), &updatedOrderStatus)
	if err != nil {
		return err
	}

	if updatedOrderStatus.OrderID == "" {
		logger.Error(domain.ErrOrderIsMandatory, "messageID", *message.MessageId)
		return domain.NewValidationError(errors.New(domain.ErrOrderIsMandatory))
	}

	if updatedOrderStatus.Status == "" {
		logger.Error(domain.ErrStatusIsMandatory, "messageID", *message.MessageId)
		return domain.NewValidationError(errors.New(domain.ErrStatusIsMandatory))
	}

	// call your use case to update the order status

	logger.Info("Message processed successfully", "orderID", updatedOrderStatus.OrderID, "status", updatedOrderStatus.Status, "staffID", updatedOrderStatus.StaffID)

	return nil // Return nil if processing is successful
}
