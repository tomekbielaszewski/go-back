package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goback",
	Short: "Goback is a cheap file archiving tool",
	Long:  `Goback is a cheap file archiving tool backed up by AWS Glacier`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(readConfig)

}

func readConfig() {

}
