package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
	"go-back/app"
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
			Bucket: aws.String(app.Config.Bucket),
		})
		if err != nil {
			log.Fatal(err)
		}

		var archives = make([]*app.Archive, len(objects.Contents))
		for i, object := range objects.Contents {
			archives[i] = &app.Archive{
				Id:     *object.ETag,
				Path:   *object.Key,
				Bucket: app.Config.Bucket,
				Size:   object.Size,
			}
		}

		var printer = &app.PrettyArchivePrinter{}
		fmt.Println(printer.Print(archives))

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
