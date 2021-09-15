package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

func init() {
	listCmd.Flags().BoolVarP(&forceUpdate, "force-update", "f", false,
		"Forces archive list update from live Glacier.")
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list backed up archives",
	Run: func(cmd *cobra.Command, args []string) {
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
