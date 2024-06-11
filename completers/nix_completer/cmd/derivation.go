package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var derivationCmd = &cobra.Command{
	Use:     "derivation",
	Short:   "work with derivations",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(derivationCmd).Standalone()

	rootCmd.AddCommand(derivationCmd)
}
