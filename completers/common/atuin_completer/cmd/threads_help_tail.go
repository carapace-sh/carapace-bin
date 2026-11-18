package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_help_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Follow threads and messages as they are recorded (until interrupted)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_help_tailCmd).Standalone()

	threads_helpCmd.AddCommand(threads_help_tailCmd)
}
