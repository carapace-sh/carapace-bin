package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_catCmd = &cobra.Command{
	Use:   "cat",
	Short: "print the contents of a file in the Nix store on stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_catCmd).Standalone()

	storeCmd.AddCommand(store_catCmd)

	// TODO positional completion
}
