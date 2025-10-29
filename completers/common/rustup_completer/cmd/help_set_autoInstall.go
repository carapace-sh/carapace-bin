package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_set_autoInstallCmd = &cobra.Command{
	Use:   "auto-install",
	Short: "The auto toolchain install mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_set_autoInstallCmd).Standalone()

	help_setCmd.AddCommand(help_set_autoInstallCmd)
}
