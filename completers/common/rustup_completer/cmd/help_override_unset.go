package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_override_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove the override toolchain for a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_override_unsetCmd).Standalone()

	help_overrideCmd.AddCommand(help_override_unsetCmd)
}
