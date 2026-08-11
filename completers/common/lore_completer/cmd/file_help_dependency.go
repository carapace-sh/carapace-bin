package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Manage file dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_dependencyCmd).Standalone()

	file_helpCmd.AddCommand(file_help_dependencyCmd)
}
