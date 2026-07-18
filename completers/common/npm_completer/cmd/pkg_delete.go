package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a value from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_deleteCmd).Standalone()

	pkgCmd.AddCommand(pkg_deleteCmd)
}
