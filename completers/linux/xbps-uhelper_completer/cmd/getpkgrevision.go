package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getpkgrevisionCommand = &cobra.Command{
	Use:   "getpkgrevisionCommand <pkgver...>",
	Short: "Prints revision from pkgvers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getpkgrevisionCommand).Standalone()

	rootCmd.AddCommand(getpkgrevisionCommand)
}
