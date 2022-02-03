package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Client) HeadObject() *s3.HeadObjectOutput {

	svc := s.Initialize()

	output, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(s.FileName),
	})

	if err != nil {
		return nil
	}
	return output
}
