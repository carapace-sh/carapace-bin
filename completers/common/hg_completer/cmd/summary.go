package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:     "summary",
	Short:   "summarize working directory state",
	Aliases: []string{"sum"},
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(summaryCmd).Standalone()

	summaryCmd.Flags().Bool("remote", false, "check for push and pull")
	rootCmd.AddCommand(summaryCmd)
}
