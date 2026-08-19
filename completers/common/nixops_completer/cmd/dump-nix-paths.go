package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var DumpNixPathsCmd = &cobra.Command{
	Use:   "dump-nix-paths",
	Short: "Dump Nix Paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(DumpNixPathsCmd).Standalone()
	rootCmd.AddCommand(DumpNixPathsCmd)
}
