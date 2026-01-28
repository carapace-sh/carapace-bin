package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var gendsaCmd = &cobra.Command{
	Use:     "gendsa",
	Short:   "Generation of DSA Private Key from Parameters",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gendsaCmd).Standalone()

	gendsaCmd.Flags().StringS("out", "out", "", "Output the key to the specified file")
	gendsaCmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	gendsaCmd.Flags().BoolS("quiet", "quiet", false, "Terse output")
	gendsaCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output")
	common.AddProviderFlags(gendsaCmd)
	common.AddRandomStateFlags(gendsaCmd)
	rootCmd.AddCommand(gendsaCmd)

	carapace.Gen(gendsaCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
	})

	carapace.Gen(gendsaCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
