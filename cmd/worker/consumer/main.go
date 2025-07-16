package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/adapter/gateway"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/core/domain"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/core/dto"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/core/port"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/core/usecase"
	appConfig "github.com/FIAP-SOAT-G20/tc4-order-service/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/infrastructure/database"
	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/infrastructure/datasource"
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

	ctx := context.Background()

	// Load AWS Config
	appCfg := appConfig.LoadConfig()

	loggerInstance := logger.NewLogger(appCfg.Environment)

	db, err := database.NewPostgresConnection(appCfg, loggerInstance)
	if err != nil {
		loggerInstance.Error("failed to connect to database", "error", err.Error())
		os.Exit(1)
	}

	orderDS := datasource.NewOrderDataSource(db.DB)
	orderHistoryDS := datasource.NewOrderHistoryDataSource(db.DB)
	orderGateway := gateway.NewOrderGateway(orderDS)
	orderHistoryGateway := gateway.NewOrderHistoryGateway(orderHistoryDS)
	orderHistoryUC := usecase.NewOrderHistoryUseCase(orderHistoryGateway)
	orderUC := usecase.NewOrderUseCase(orderGateway, orderHistoryUC)

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

			go processJob(ctx, message, sqsClient, appCfg.AWS_SQSURL, loggerInstance, orderUC)
		}
		wg.Wait()
	}
}

func processJob(
	ctx context.Context,
	message types.Message,
	sqsClient *sqs.Client,
	awsSQSURL string,
	logger *logger.Logger,
	uc port.OrderUseCase,
) {
	defer wg.Done()

	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(awsSQSURL),
		ReceiptHandle: message.ReceiptHandle,
	}

	err := processedMessage(ctx, message, logger, uc)
	if err != nil {
		logger.Error("failed to process message", "error", err.Error(), "messageID", *message.MessageId)
		// send the message to a dead-letter queue for further investigation
	}

	_, err = sqsClient.DeleteMessage(context.TODO(), deleteParams)

	if err != nil {
		log.Fatal(err)
	}
}

func processedMessage(ctx context.Context, message types.Message, logger *logger.Logger, uc port.OrderUseCase) error {
	// Here you can implement the logic to process the message
	// For example, you can unmarshal the message body and update the order status in your database
	logger.Info("Processing message", "messageID", *message.MessageId, "body", *message.Body)

	// Unmarshal the message body to your entity
	var updatedOrderStatus entity.OrderStatusUpdated
	err := json.Unmarshal([]byte(*message.Body), &updatedOrderStatus)
	if err != nil {
		return err
	}

	if updatedOrderStatus.OrderID == 0 {
		logger.Error(domain.ErrOrderIsMandatory, "messageID", *message.MessageId)
		return domain.NewValidationError(errors.New(domain.ErrOrderIsMandatory))
	}

	if updatedOrderStatus.Status == "" {
		logger.Error(domain.ErrStatusIsMandatory, "messageID", *message.MessageId)
		return domain.NewValidationError(errors.New(domain.ErrStatusIsMandatory))
	}

	// Get Order by ID
	_, err = uc.Get(ctx, dto.GetOrderInput{ID: updatedOrderStatus.OrderID})
	if err != nil {
		return err
	}

	// Update the order status in the database
	uoi := dto.UpdateOrderInput{
		ID:     updatedOrderStatus.OrderID,
		Status: updatedOrderStatus.Status,
	}
	if updatedOrderStatus.StaffID != nil {
		uoi.StaffID = *updatedOrderStatus.StaffID
	}
	_, err = uc.Update(ctx, uoi)
	if err != nil {
		return err
	}

	logger.Info("Message processed successfully", "orderID", updatedOrderStatus.OrderID, "status", updatedOrderStatus.Status, "staffID", updatedOrderStatus.StaffID)

	return nil // Return nil if processing is successful
}
