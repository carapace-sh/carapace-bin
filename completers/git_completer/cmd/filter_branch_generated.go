package cmd

import (
	"github.com/spf13/cobra"
)

var filter_branchCmd = &cobra.Command{
	Use:     "filter-branch",
	Short:   "Rewrite branches",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {

	rootCmd.AddCommand(filter_branchCmd)
}
