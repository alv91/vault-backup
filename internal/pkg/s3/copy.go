package s3

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Client) CopyObject(fileName string) error {

	svc := s.Initialize()

	source := s.Bucket + "/" + fileName

	_, err := svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(s.Bucket),
		CopySource: aws.String(url.PathEscape(source)),
		Key:        aws.String(s.FileName),
	})
	if err != nil {
		fmt.Errorf("Unable to copy file '%s' from bucket %s to %s, %v", fileName, s.Bucket, s.FileName, err)

		return err
	}

	return nil
}
