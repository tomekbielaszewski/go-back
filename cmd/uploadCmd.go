package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
	"log"
)

type UploadCmdArgs struct {
	file string
}

func init() {
	uploadCmd.Flags().StringVarP(&uploadCmdArgs.file, "file", "f", "", "Path to archive")
	rootCmd.AddCommand(uploadCmd)
}

var uploadCmdArgs = &UploadCmdArgs{
	file: "",
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "uploads the archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uploadCmdArgs)

		defaultConfig, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		s3Client := s3.NewFromConfig(defaultConfig)

		objects, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String("some-bucket"),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Printing stuff:")
		for _, object := range objects.Contents {
			log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
		}
	},
}
