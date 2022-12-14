package cmd

import (
	"github.com/spf13/cobra"
)

var mergetoolCmd = &cobra.Command{
	Use:     "mergetool",
	Short:   "Run merge conflict resolution tools to resolve merge conflicts",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {

	rootCmd.AddCommand(mergetoolCmd)
}
