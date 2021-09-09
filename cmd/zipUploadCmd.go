package cmd

import "github.com/spf13/cobra"

var zipUploadCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(zipUploadCmd)
}
