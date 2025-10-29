package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolchain_help_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Create a custom toolchain by symlinking to a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_help_linkCmd).Standalone()

	toolchain_helpCmd.AddCommand(toolchain_help_linkCmd)
}
