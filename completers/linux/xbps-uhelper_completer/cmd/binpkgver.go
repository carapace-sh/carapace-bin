package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var binpkgverCommand = &cobra.Command{
	Use:   "binpkgver <binpkg...>",
	Short: "Prints the pkgver of binpkg names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(binpkgverCommand).Standalone()

	rootCmd.AddCommand(binpkgverCommand)

	carapace.Gen(binpkgverCommand).PositionalCompletion(carapace.ActionFiles(".xbps"))
}
