package utils

import (
	"context"
	"log"
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
			Key:    aws.String(name),
		}, func(opts *s3.PresignOptions) {
			opts.Expires = expiry
		})
	if err != nil {
		return "", nil, err
	}

	headers := map[string]string{}

	for k, v := range request.SignedHeader {
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
			Key:    aws.String(name),
		}, func(opts *s3.PresignOptions) {
			opts.Expires = expiry
		})
	if err != nil {
		return "", nil, err
	}

	headers := map[string]string{}

	for k, v := range request.SignedHeader {
		headers[k] = v[0]
	}

	return request.URL, headers, nil
}

func (awsMgr *AWS) GetSPBucketSize(bucket, userID string) (uint64, int) {

	prefix := "sp_data_" + userID

	maxKeys := int(1000)

	cl := s3.NewFromConfig(awsMgr.Config)

	// Set the parameters based on the CLI flag inputs.
	params := &s3.ListObjectsV2Input{
		Bucket: &bucket,
	}

	if len(prefix) != 0 {
		params.Prefix = &prefix
	}

	// Create the Paginator for the ListObjectsV2 operation.
	p := s3.NewListObjectsV2Paginator(cl, params, func(o *s3.ListObjectsV2PaginatorOptions) {
		if v := int32(maxKeys); v != 0 {
			o.Limit = v
		}
	})

	// Iterate through the S3 object pages, printing each object returned.
	var i int
	totalSize := uint64(0)
	itemCount := int(0)
	for p.HasMorePages() {
		i++

		// Next Page takes a new context for each page retrieval. This is where
		// you could add timeouts or deadlines.
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("failed to get page %v, %v", i, err)
		}

		// Log the objects found
		for _, obj := range page.Contents {
			totalSize += uint64(obj.Size)
			itemCount += 1
		}
	}

	return totalSize, itemCount
}
