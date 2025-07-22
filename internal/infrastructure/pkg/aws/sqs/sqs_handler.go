package sqs

import (
	"context"
	"fmt"

	"github.com/FIAP-SOAT-G20/tc4-order-service/internal/infrastructure/logger"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SqsHandler struct {
	sqsClient       *SqsClient
	queueURL        string
	maxMessages     int
	waitTimeSeconds int
	logger          *logger.Logger
}

// NewSqsHandler creates a new SQS handler with the provided SQS client
func NewSqsHandler(sqsClient *SqsClient, queueURL string, maxMessages int, waitTimeSeconds int, logger *logger.Logger) *SqsHandler {
	return &SqsHandler{
		sqsClient:       sqsClient,
		logger:          logger,
		queueURL:        queueURL,
		maxMessages:     maxMessages,
		waitTimeSeconds: waitTimeSeconds,
	}
}

// SendMessage sends a single message to the configured queue
func (h *SqsHandler) SendMessage(ctx context.Context, messageBody string) (*types.Message, error) {
	h.logger.Info("Sending message to queue", "queueURL", h.queueURL, "messageBody", messageBody)

	message, err := h.sqsClient.SendMessage(ctx, h.queueURL, messageBody)
	if err != nil {
		h.logger.Error("Failed to send message to queue",
			"error", err.Error(),
			"queueURL", h.queueURL,
			"messageBody", messageBody,
		)
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	h.logger.Info("Successfully sent message to queue",
		"messageId", *message.MessageId,
		"queueURL", h.queueURL,
	)

	return message, nil
}

// ReceiveMessages processes messages and optionally deletes them after processing
func (h *SqsHandler) ReceiveMessages(ctx context.Context, processor func(types.Message) (bool, error)) error {
	messages, err := h.sqsClient.ReceiveMessages(ctx, h.queueURL, h.maxMessages, h.waitTimeSeconds)
	if err != nil {
		return fmt.Errorf("failed to receive messages: %w", err)
	}

	for _, message := range messages {
		if reprocess, err := processor(message); err != nil {
			h.logger.Error("Failed to process message",
				"error", err.Error(),
				"messageId", *message.MessageId,
				"messageBody", *message.Body,
			)

			if reprocess {
				h.logger.Info("Reprocessing message", "messageId", *message.MessageId)
				continue // Reprocess the message
			}

			h.logger.Info("Sending message to dead-letter queue", "messageId", *message.MessageId)
			// TODO: implement dead-letter queue logic
		}

		if err := h.DeleteMessage(ctx, h.queueURL, *message.ReceiptHandle); err != nil {
			h.logger.Error("Failed to delete message",
				"error", err.Error(),
				"messageId", *message.MessageId,
				"messageBody", *message.Body,
			)
			continue
		}
	}

	return nil
}

// DeleteMessage deletes a specific message from the queue
func (h *SqsHandler) DeleteMessage(ctx context.Context, queueURL string, receiptHandle string) error {
	return h.sqsClient.DeleteMessage(ctx, queueURL, receiptHandle)
}
