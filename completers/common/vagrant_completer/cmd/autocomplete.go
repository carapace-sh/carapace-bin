package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "manages autocomplete installation on host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autocompleteCmd).Standalone()

	rootCmd.AddCommand(autocompleteCmd)
}
