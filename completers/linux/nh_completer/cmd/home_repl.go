package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var home_replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Load a home-manager configuration in a Nix REPL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(home_replCmd).Standalone()

	home_replCmd.Flags().StringP("configuration", "c", "", "Name of the flake homeConfigurations attribute")

	homeCmd.AddCommand(home_replCmd)
}
