package cmd

import (
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:     "stash",
	Short:   "Stash the changes in a dirty working directory away",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	rootCmd.AddCommand(stashCmd)
}
