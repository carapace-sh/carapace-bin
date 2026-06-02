package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var darwin_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Build and activate a nix-darwin configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(darwin_switchCmd).Standalone()

	addDarwinRebuildFlags(darwin_switchCmd)

	darwinCmd.AddCommand(darwin_switchCmd)
}
