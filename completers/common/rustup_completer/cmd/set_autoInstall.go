package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_autoInstallCmd = &cobra.Command{
	Use:   "auto-install",
	Short: "The auto toolchain install mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_autoInstallCmd).Standalone()

	set_autoInstallCmd.Flags().BoolP("help", "h", false, "Print help")
	setCmd.AddCommand(set_autoInstallCmd)
}
