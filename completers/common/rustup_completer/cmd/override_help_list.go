package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var override_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List directory toolchain overrides",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(override_help_listCmd).Standalone()

	override_helpCmd.AddCommand(override_help_listCmd)
}
