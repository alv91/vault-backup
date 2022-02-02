package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Client struct {
	AccessKey       string
	SecretAccessKey string
	Region          string
	Bucket          string
	Endpoint        string
	FileName        string
}

func NewClient(accessKey, secretKey, region, bucket, endpoint, fileName string) *Client {
	return &Client{
		accessKey,
		secretKey,
		region,
		bucket,
		endpoint,
		fileName,
	}
}

func (s *Client) Initialize() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s.AccessKey, s.SecretAccessKey, ""),
		Region:      aws.String(s.Region),
		Endpoint:    aws.String(s.Endpoint),
	},
	)
	if err != nil {
		fmt.Println("Cant establish S3 connection, err: ", err)
	}

	return s3.New(sess)
}
