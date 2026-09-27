package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pkg_fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Auto corrects common errors in package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pkg_fixCmd).Standalone()

	help_pkgCmd.AddCommand(help_pkg_fixCmd)
}
