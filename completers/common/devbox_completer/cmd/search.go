package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <pkg>",
	Short: "Search for nix packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("show-all", false, "show all available templates")
	rootCmd.AddCommand(searchCmd)

	// TODO positional completion
}
