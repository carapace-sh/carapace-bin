package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_addFileCmd = &cobra.Command{
	Use:   "add-file",
	Short: "add a regular file to the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_addFileCmd).Standalone()

	store_addFileCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	store_addFileCmd.Flags().BoolP("name", "n", false, "Override the name component of the store path")
	storeCmd.AddCommand(store_addFileCmd)

	addLoggingFlags(store_addFileCmd)

	carapace.Gen(store_addFileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
