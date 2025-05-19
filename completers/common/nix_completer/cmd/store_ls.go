package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "show information about a path in the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_lsCmd).Standalone()

	store_lsCmd.Flags().BoolP("directory", "d", false, "Show directories rather than their contents")
	store_lsCmd.Flags().Bool("json", false, "Produce output in JSON format")
	store_lsCmd.Flags().BoolP("long", "l", false, "Show detailed file information")
	store_lsCmd.Flags().BoolP("recursive", "R", false, "List subdirectories recursively")
	storeCmd.AddCommand(store_lsCmd)

	addLoggingFlags(store_lsCmd)

	carapace.Gen(store_lsCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
