package cmd

import (
	"github.com/spf13/cobra"
)

var show_indexCmd = &cobra.Command{
	Use:     "show-index",
	Short:   "Show packed archive index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {

	rootCmd.AddCommand(show_indexCmd)
}
