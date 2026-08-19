package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpNixPathsCmd = &cobra.Command{
	Use:   "dump-nix-paths",
	Short: "dump Nix paths referenced in deployments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpNixPathsCmd).Standalone()
	rootCmd.AddCommand(dumpNixPathsCmd)
}
