package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_threads_resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Pick a captured thread and resume it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_threads_resumeCmd).Standalone()

	help_threadsCmd.AddCommand(help_threads_resumeCmd)
}
