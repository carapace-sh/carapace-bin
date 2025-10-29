package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var override_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the override toolchain for a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(override_help_setCmd).Standalone()

	override_helpCmd.AddCommand(override_help_setCmd)
}
