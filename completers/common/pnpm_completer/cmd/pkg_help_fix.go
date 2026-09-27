package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_help_fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Auto corrects common errors in package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_help_fixCmd).Standalone()

	pkg_helpCmd.AddCommand(pkg_help_fixCmd)
}
