package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_help_acquireCmd = &cobra.Command{
	Use:   "acquire",
	Short: "Acquire lock on file(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_help_acquireCmd).Standalone()

	lock_helpCmd.AddCommand(lock_help_acquireCmd)
}
