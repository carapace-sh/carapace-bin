package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getnameCommand = &cobra.Command{
	Use:   "getname <pkgver|dep...>",
	Short: "Prints pkgname from pkgvers or dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getnameCommand).Standalone()

	rootCmd.AddCommand(getnameCommand)
}
