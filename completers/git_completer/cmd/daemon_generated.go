package cmd

import (
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:     "daemon",
	Short:   "A really simple server for Git repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {

	rootCmd.AddCommand(daemonCmd)
}
