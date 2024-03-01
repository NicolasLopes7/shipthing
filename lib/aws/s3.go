package aws

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/NicolasLopes7/shipthing/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"golang.org/x/sync/errgroup"
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

func DownloadS3Folder(deployId string) error {
	files, err := config.S3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Prefix: aws.String(fmt.Sprintf("tmp/%s", deployId)),
	})

	if err != nil {
		return err
	}

	group, _ := errgroup.WithContext(context.Background())
	for _, c := range files.Contents {
		c := c
		group.Go(func() error {
			if c.Key == nil {
				return nil
			}

			dirname, err := os.Getwd()
			if err != nil {
				return err
			}

			outputPath := path.Join(dirname, *c.Key)
			dirName := path.Dir(outputPath)

			if _, err := os.Stat(dirName); os.IsNotExist(err) {
				err := os.MkdirAll(dirName, 0755)
				if err != nil {
					return err
				}
			}

			res, err := config.S3Client.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
				Key:    c.Key,
			})

			if err != nil {
				return err
			}

			file, err := os.Create(outputPath)
			if err != nil {
				return err
			}

			defer file.Close()

			_, err = file.ReadFrom(res.Body)

			if err != nil {
				return err
			}

			defer res.Body.Close()

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
