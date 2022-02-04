package s3

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const S3_ACL = "private"

func (s *Client) PutObject(rs io.ReadSeeker, fileName string) error {
	svc := s.Initialize()

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileName),
		ACL:    aws.String(S3_ACL),
		Body:   rs,
	})
	if err != nil {
		fmt.Errorf("cant upload file with name %s, err: %e", fileName, err)

		return err
	}

	return nil
}
