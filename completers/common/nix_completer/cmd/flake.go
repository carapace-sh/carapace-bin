package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flakeCmd = &cobra.Command{
	Use:     "flake",
	Short:   "manage Nix flakes",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flakeCmd).Standalone()

	rootCmd.AddCommand(flakeCmd)
}
