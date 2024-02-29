package aws

import (
	"os"

	"github.com/NicolasLopes7/shipthing/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadToS3(path string) error {
	_, err := config.S3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(path),
		Body:   os.Stdin,
	})

	if err != nil {
		return err
	}

	return nil
}
