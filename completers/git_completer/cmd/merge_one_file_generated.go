package cmd

import (
	"github.com/spf13/cobra"
)

var merge_one_fileCmd = &cobra.Command{
	Use:     "merge-one-file",
	Short:   "The standard helper program to use with git-merge-index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {

	rootCmd.AddCommand(merge_one_fileCmd)
}
