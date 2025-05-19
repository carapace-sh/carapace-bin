package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_addPathCmd = &cobra.Command{
	Use:   "add-path",
	Short: "add a path to the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_addPathCmd).Standalone()

	store_addPathCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	store_addPathCmd.Flags().BoolP("name", "n", false, "Override the name component of the store path")
	storeCmd.AddCommand(store_addPathCmd)

	addLoggingFlags(store_addPathCmd)

	// TODO positional completion
}
