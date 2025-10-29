package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_help_autoInstallCmd = &cobra.Command{
	Use:   "auto-install",
	Short: "The auto toolchain install mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_help_autoInstallCmd).Standalone()

	set_helpCmd.AddCommand(set_help_autoInstallCmd)
}
