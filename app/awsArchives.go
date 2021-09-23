package app

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func DownloadArchiveList(gobackConfig *GobackConfig) []*Archive {
	defaultConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(defaultConfig)

	objects, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(gobackConfig.Bucket),
	})
	if err != nil {
		log.Fatal(err)
	}

	var archives = make([]*Archive, len(objects.Contents))
	for i, object := range objects.Contents {
		archives[i] = &Archive{
			Id:     *object.ETag,
			Path:   *object.Key,
			Bucket: gobackConfig.Bucket,
			Size:   object.Size,
		}
	}
	return archives
}
