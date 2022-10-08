package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new package to your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()
	rootCmd.AddCommand(addCmd)

	// TODO nix package search
}
