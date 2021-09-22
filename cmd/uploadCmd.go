package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	},
}
