package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_threads_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Follow threads and messages as they are recorded (until interrupted)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_threads_tailCmd).Standalone()

	help_threadsCmd.AddCommand(help_threads_tailCmd)
}
