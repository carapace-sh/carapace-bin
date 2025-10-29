package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolchain_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_help_listCmd).Standalone()

	toolchain_helpCmd.AddCommand(toolchain_help_listCmd)
}
