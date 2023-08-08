package utils

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWS struct {
	Config aws.Config

}

func NewAWS() (*AWS, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return &AWS{
		Config: cfg,
	}, nil
}

func (awsMgr *AWS) CreateSignedPutUrl(bucket, name string,
	expiry time.Duration) (string, map[string]string, error) {
	cl := s3.NewFromConfig(awsMgr.Config)
	s3Cl := s3.NewPresignClient(cl)
	request, err := s3Cl.PresignPutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key: aws.String(name),
		}, func(opts *s3.PresignOptions) {
			opts.Expires = expiry
		})
	if (err != nil) {
		return "", nil, err
	}

	headers := map[string]string{}

	for k,v := range request.SignedHeader {
		headers[k] = v[0]
	}

	return request.URL, headers, nil
}

func (awsMgr *AWS) CreateSignedGetUrl(bucket, name string,
	expiry time.Duration) (string, map[string]string, error) {
	cl := s3.NewFromConfig(awsMgr.Config)
	s3Cl := s3.NewPresignClient(cl)
	request, err := s3Cl.PresignGetObject(context.TODO(),
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key: aws.String(name),
		}, func(opts *s3.PresignOptions) {
			opts.Expires = expiry
		})
	if (err != nil) {
		return "", nil, err
	}

	headers := map[string]string{}

	for k,v := range request.SignedHeader {
		headers[k] = v[0]
	}

	return request.URL, headers, nil
}
