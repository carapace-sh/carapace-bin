package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_override_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the override toolchain for a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_override_setCmd).Standalone()

	help_overrideCmd.AddCommand(help_override_setCmd)
}
