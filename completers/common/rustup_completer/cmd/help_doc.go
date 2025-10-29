package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Open the documentation for the current toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_docCmd).Standalone()

	helpCmd.AddCommand(help_docCmd)
}
