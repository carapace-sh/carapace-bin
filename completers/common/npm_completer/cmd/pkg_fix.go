package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix package.json fields",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_fixCmd).Standalone()

	pkgCmd.AddCommand(pkg_fixCmd)
}
