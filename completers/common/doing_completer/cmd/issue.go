package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue [OPTIONS] COMMAND [ARGS]...",
	Short: "Work with issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issueCmd).Standalone()

	issueCmd.Flags().Bool("help", false, "Show this message and exit.")
	rootCmd.AddCommand(issueCmd)
}
