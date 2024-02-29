package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

var (
	AwsSession *session.Session
	S3Uploader *s3manager.Uploader
)

func InitConfig() error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	AwsSession, err := createAwsSession()

	if err != nil {
		return err
	}

	S3Uploader = createS3Uploader(AwsSession)

	return nil
}

func createAwsSession() (*session.Session, error) {
	sesh, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})

	if err != nil {
		return nil, err
	}

	return sesh, nil
}

func createS3Uploader(session *session.Session) *s3manager.Uploader {
	return s3manager.NewUploader(session)
}
