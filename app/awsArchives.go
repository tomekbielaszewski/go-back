package app

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"log"
	"strings"
)

func DownloadArchiveList(gobackConfig *GobackConfig) ([]*Archive, error) {
	s3Client, err := createS3Client()
	if err != nil {
		log.Fatal(err)
	}

	objects, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(gobackConfig.Bucket),
	})
	if err != nil {
		if opErr, ok := err.(*smithy.OperationError); ok {
			if strings.Contains(opErr.Err.Error(), "NoSuchBucket") {
				return recoverFromNoBucketError(gobackConfig), nil
			}
		} else {
			log.Fatal(err)
			return nil, err
		}
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
	return archives, nil
}

func createS3Client() (*s3.Client, error) {
	defaultConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(defaultConfig)
	return s3Client, nil
}

func recoverFromNoBucketError(gobackConfig *GobackConfig) []*Archive {
	switch gobackConfig.NoBucketAction {
	case CREATE:
		fmt.Printf("There is no bucket \"%s\" you were looking for. Let me create it for you!\n", gobackConfig.Bucket)
		createBucket(gobackConfig.Bucket)
	case ASK:
		fmt.Printf("There is no bucket \"%s\" you were looking for. Do you wish me to create it for you?\n", gobackConfig.Bucket)
		askToCreateBucket(gobackConfig.Bucket)
	case EXIT:
		fmt.Printf("There is no bucket \"%s\" you were looking for. Create it first!\n", gobackConfig.Bucket)
	}

	return []*Archive{}
}

func askToCreateBucket(bucket string) {
	var input string
	for {
		fmt.Print("yes or no: ")
		fmt.Scanf("%s", &input)

		input = strings.ToLower(input)

		if input == "yes" || input == "y" {
			createBucket(bucket)
			break
		}
		if input == "no" || input == "n" {
			fmt.Println("Ok! Bye!")
			break
		}
	}
}

func createBucket(bucket string) {
	fmt.Printf("Trrrr.... creating bucket \"%s\".... trrrr", bucket)
}
