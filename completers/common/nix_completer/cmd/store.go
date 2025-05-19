package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:     "store",
	Short:   "manipulate a Nix store",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storeCmd).Standalone()

	rootCmd.AddCommand(storeCmd)
}
