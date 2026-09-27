package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pkg_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a key from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pkg_deleteCmd).Standalone()

	help_pkgCmd.AddCommand(help_pkg_deleteCmd)
}
