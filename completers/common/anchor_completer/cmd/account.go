package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Fetch and deserialize an account using the IDL provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountCmd).Standalone()

	accountCmd.Flags().BoolP("help", "h", false, "Print help")
	accountCmd.Flags().String("idl", "", "Path of IDL to use (defaults to workspace IDL)")
	rootCmd.AddCommand(accountCmd)

	carapace.Gen(accountCmd).FlagCompletion(carapace.ActionMap{
		"idl": carapace.ActionFiles(".json"),
	})

	carapace.Gen(accountCmd).PositionalCompletion(
		carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return anchor.ActionPrograms().Suffix(".")
			}
			return carapace.ActionValues()
		}),
	)
}
