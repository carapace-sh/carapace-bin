package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var darwin_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a nix-darwin configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(darwin_buildCmd).Standalone()

	addDarwinRebuildFlags(darwin_buildCmd)

	darwinCmd.AddCommand(darwin_buildCmd)
}
