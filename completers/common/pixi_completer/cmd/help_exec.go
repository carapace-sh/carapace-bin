package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command and install it in a temporary environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_execCmd).Standalone()

	helpCmd.AddCommand(help_execCmd)
}
