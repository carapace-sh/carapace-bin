package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var realisationCmd = &cobra.Command{
	Use:     "realisation",
	Short:   "manipulate a Nix realisation",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(realisationCmd).Standalone()

	rootCmd.AddCommand(realisationCmd)
}
