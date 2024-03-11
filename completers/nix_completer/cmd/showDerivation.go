package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showDerivationCmd = &cobra.Command{
	Use:     "show-derivation",
	Short:   "show the contents of a store derivation",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showDerivationCmd).Standalone()

	showDerivationCmd.Flags().BoolP("recursive", "r", false, "Include the dependencies of the specified derivations")
	rootCmd.AddCommand(showDerivationCmd)

	addEvaluationFlags(showDerivationCmd)
	addFlakeFlags(showDerivationCmd)
	addLoggingFlags(showDerivationCmd)

	// TODO positional completion
}
