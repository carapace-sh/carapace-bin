package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a key from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_deleteCmd).Standalone()

	pkg_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pkgCmd.AddCommand(pkg_deleteCmd)
}
