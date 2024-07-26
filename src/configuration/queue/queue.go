package queue

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	AWS_REGION = "AWS_REGION"
	AWS_HOST   = "AWS_HOST"
)

func initAWSSession() *session.Session {

	awsRegion := os.Getenv(AWS_REGION)
	awsHost := os.Getenv(AWS_HOST)

	return session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(awsHost),
		Region:   aws.String(awsRegion),
	}))
}

func GetSQSClient() *sqs.SQS {
	sess := initAWSSession()
	return sqs.New(sess)
}
