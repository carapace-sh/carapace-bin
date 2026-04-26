package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var feedbackCmd = &cobra.Command{
	Use:   "feedback [title]",
	Short: "Submit feedback for gh-stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(feedbackCmd).Standalone()

	rootCmd.AddCommand(feedbackCmd)
}
