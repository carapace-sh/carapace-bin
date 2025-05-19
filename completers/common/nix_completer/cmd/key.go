package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:     "key",
	Short:   "generate and convert Nix signing keys",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyCmd).Standalone()

	rootCmd.AddCommand(keyCmd)
}
