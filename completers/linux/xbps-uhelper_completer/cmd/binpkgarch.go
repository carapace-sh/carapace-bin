package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var binpkgarchCommand = &cobra.Command{
	Use:   "binpkgarch <binpkg>",
	Short: "Prints the architecture of binpkg names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(binpkgarchCommand).Standalone()

	rootCmd.AddCommand(binpkgarchCommand)

	carapace.Gen(binpkgarchCommand).PositionalCompletion(carapace.ActionFiles(".xbps"))
}
