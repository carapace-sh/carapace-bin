package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_help_resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Pick a captured thread and resume it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_help_resumeCmd).Standalone()

	threads_helpCmd.AddCommand(threads_help_resumeCmd)
}
