package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpStoresCmd = &cobra.Command{
	Use:     "help",
	Short:   "show help about store types and their settings",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpStoresCmd).Standalone()

	rootCmd.AddCommand(helpStoresCmd)

	// TODO positional completion
}
