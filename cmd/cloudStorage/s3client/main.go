package s3client

import (
	"context"
	"io"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var (
	s3Client *s3.Client
	uploader *manager.Uploader
	once     sync.Once
)

func setup() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	accessKey := os.Getenv("AWS_KEY")
	secretKey := os.Getenv("AWS_SECRET")
	endpoint := os.Getenv("AWS_ENDPOINT")

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               endpoint,
			HostnameImmutable: true,
			Source:            aws.EndpointSourceCustom,
		}, nil
	})
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithRegion("auto"),
	)

	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	s3Client = s3.NewFromConfig(awsConfig) // Initialize S3 client
}

func UploadFile(ctx context.Context, bucket, filename string, file io.Reader) (string, error) {
	once.Do(setup)

	uploader = manager.NewUploader(s3Client)

	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    "public-read",
	})

	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}

	return result.Location, nil
}
