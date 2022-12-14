package cmd

import (
	"github.com/spf13/cobra"
)

var fetch_packCmd = &cobra.Command{
	Use:     "fetch-pack",
	Short:   "Receive missing objects from another repository",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {

	rootCmd.AddCommand(fetch_packCmd)
}
