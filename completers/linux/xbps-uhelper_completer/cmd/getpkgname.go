package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getpkgnameCommand = &cobra.Command{
	Use:   "getpkgname <pkgver...>",
	Short: "Prints pkgname from pkgvers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getpkgnameCommand).Standalone()

	rootCmd.AddCommand(getpkgnameCommand)
}
