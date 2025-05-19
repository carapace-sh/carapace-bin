package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mergeOneFileCmd = &cobra.Command{
	Use:     "merge-one-file",
	Short:   "The standard helper program to use with git-merge-index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(mergeOneFileCmd).Standalone()

	rootCmd.AddCommand(mergeOneFileCmd)
}
