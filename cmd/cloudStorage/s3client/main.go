package s3client

import (
	"context"
	"fmt"
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

	"github.com/cheggaaa/pb/v3"
)

var (
	s3Client   *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
	once       sync.Once
)

type progressWriter struct {
	writer  io.WriterAt
	size    int64
	bar     *pb.ProgressBar
	display bool
}

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

func DownloadFile(ctx context.Context, bucket, key string, file io.WriterAt, displayProgressBar bool) (string, error) {
	once.Do(setup)

	downloader = manager.NewDownloader(s3Client)

	// Get the object size
	s3ObjectSize := GetS3ObjectSize(bucket, key)

	fmt.Println("total content size", s3ObjectSize)

	// Initialize progress writer
	writer := &progressWriter{writer: file, size: s3ObjectSize}
	writer.display = displayProgressBar
	writer.init(s3ObjectSize)

	// Start the download
	numBytes, err := downloader.Download(ctx, file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("Failed to download "+key+" from s3 bucket "+bucket, fmt.Errorf("Error downloading from file", err))
	}
	writer.finish()

	fmt.Println("Download completed", file, numBytes, "bytes")

	if err != nil {
		log.Fatalf("Failed to download file, %v", err)
	}

	return "", nil
}

func GetS3ObjectSize(bucket, item string) int64 {

	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	}

	result, err := s3Client.HeadObject(context.Background(), input)
	if err != nil {
		log.Fatalf("Error getting size of file, %v", err)
	}

	return *result.ContentLength
}

func (pw *progressWriter) init(s3ObjectSize int64) {
	if pw.display {
		pw.bar = pb.StartNew(int(s3ObjectSize))
		pw.bar.Set(pb.Bytes, true)
	}
}

func (pw *progressWriter) WriteAt(p []byte, off int64) (int, error) {
	if pw.display {
		pw.bar.Add64(int64(len(p)))
	}
	return pw.writer.WriteAt(p, off)
}

func (pw *progressWriter) finish() {
	if pw.display {
		pw.bar.Finish()
	}
}
