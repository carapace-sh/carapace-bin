package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a shell command in the context of a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_execCmd).Standalone()

	helpCmd.AddCommand(help_execCmd)
}
