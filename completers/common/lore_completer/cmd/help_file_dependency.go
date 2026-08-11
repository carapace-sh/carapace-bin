package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Manage file dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_dependencyCmd).Standalone()

	help_fileCmd.AddCommand(help_file_dependencyCmd)
}
