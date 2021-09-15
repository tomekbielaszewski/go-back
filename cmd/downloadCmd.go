package cmd

import "github.com/spf13/cobra"

var shortId string

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	downloadCmd.Flags().StringVarP(&shortId, "short-id", "i", "", "ID of backed up resource. List your available resources with \"list\" command")
	rootCmd.AddCommand(downloadCmd)
}
