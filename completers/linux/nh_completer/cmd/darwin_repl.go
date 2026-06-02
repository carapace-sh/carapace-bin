package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var darwin_replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Load a nix-darwin configuration in a Nix REPL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(darwin_replCmd).Standalone()

	darwin_replCmd.Flags().StringP("hostname", "H", "", "Select this hostname from darwinConfigurations")

	darwinCmd.AddCommand(darwin_replCmd)
}
