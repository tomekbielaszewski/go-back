package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func init() {
	listCmd.Flags().BoolVarP(&forceUpdate, "force-update", "f", false, "Forces archive list update from live Glacier.")
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list backed up archives",
	Run: func(cmd *cobra.Command, args []string) {

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

		if isUpdateForced() {
			updateArchives()
		}
		printArchives()
	},
}

func printArchives() {

}

func updateArchives() {

}

func isUpdateForced() bool {
	return forceUpdate || isArchiveStale()
}

func isArchiveStale() bool {
	return time.Now().After(lastUpdate().Add(stalePeriod()))
}

func stalePeriod() time.Duration {
	return time.Hour * 24 //todo this should ne taken from state file
}

func lastUpdate() time.Time {
	return time.Now() //todo it should be taken from state file
}

var forceUpdate bool
