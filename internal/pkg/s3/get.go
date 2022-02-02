package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Client) GetObject(filename string) *s3.GetObjectOutput {
	svc := s.Initialize()

	output, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		return nil
	}

	return output
}
