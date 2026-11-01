package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "Evaluate any attribute in devenv.nix and return JSON",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evalCmd).Standalone()

	rootCmd.AddCommand(evalCmd)
}
