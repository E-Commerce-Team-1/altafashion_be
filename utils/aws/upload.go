package aws

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// CREATE RANDOM STRING

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func autoGenerate(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return autoGenerate(length, charset)
}

func UploadProfileProduct(c echo.Context) (string, error) {

	file, fileheader, err := c.Request().FormFile("image")
	if err != nil {
		log.Print(err)
		return "", err
	}

	randomStr := String(20)

	godotenv.Load(".env")

	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_USER"), os.Getenv("AWS_KEY"), ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String("altafashion"),                                      // bucket's name
		Key:         aws.String("myfiles/" + randomStr + "-" + fileheader.Filename), // files destination location
		Body:        file,                                                           // content of the file
		ContentType: aws.String("image/jpg"),                                        // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), input)

	// RETURN URL LOCATION IN AWS
	return res.Location, err
}
