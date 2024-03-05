package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	AwsSession  *session.Session
	S3Uploader  *s3manager.Uploader
	S3Client    *s3.S3
	RedisClient *redis.Client
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
	S3Client = createS3Client(AwsSession)
	RedisClient = createRedisClient()

	return nil
}

func createAwsSession() (*session.Session, error) {
	sesh, err := session.NewSession(&aws.Config{
		Region:           aws.String(os.Getenv("AWS_REGION")),
		Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT")),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		return nil, err
	}

	return sesh, nil
}

func createS3Uploader(session *session.Session) *s3manager.Uploader {
	return s3manager.NewUploader(session)
}

func createS3Client(session *session.Session) *s3.S3 {
	return s3.New(session)
}

func createRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return redisClient
}
