package queue

import (
	"fmt"
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

const (
	AWS_REGION = "AWS_REGION"
	AWS_HOST   = "AWS_HOST"
)

var (
	queueURL = "AWS_QUEUE_URL"
)

func initAWSSession() *session.Session {

	awsRegion := os.Getenv(AWS_REGION)
	awsHost := os.Getenv(AWS_HOST)

	return session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(awsHost),
		Region:   aws.String(awsRegion),
	}))
}

func getSQSClient() *sqs.SQS {
	sess := initAWSSession()
	return sqs.New(sess)
}

func QueueSendMessage(message string) *rest_err.RestErr {
	queue := getSQSClient()
	queueURLHost := os.Getenv(queueURL)

	_, err := queue.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURLHost),
		MessageBody: aws.String(message),
	})
	if err != nil {
		logger.Error("error to send message to queue", err, zap.String("journey", "sendMessage"))
		return &rest_err.RestErr{
			Code:    500,
			Message: err.Error(),
			Err:     "error to send message to queue",
		}
	}

	return nil
}

func QueueDeleteMessage(queueMessage *sqs.DeleteMessageInput) *rest_err.RestErr {
	queue := getSQSClient()

	deleteMessageOutpout, err := queue.DeleteMessage(queueMessage)
	if err != nil {
		logger.Error("error to delete message from queue", err, zap.String("journey", "deleteMessage"))
		return &rest_err.RestErr{
			Code:    500,
			Message: err.Error(),
			Err:     "error to delete message from queue",
		}
	}

	fmt.Println(deleteMessageOutpout)
	return nil
}
