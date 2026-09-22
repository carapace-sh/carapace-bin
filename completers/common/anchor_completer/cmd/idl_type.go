package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_typeCmd = &cobra.Command{
	Use:   "type",
	Short: "Generate TypeScript type for the IDL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_typeCmd).Standalone()

	idl_typeCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_typeCmd.Flags().StringP("out", "o", "", "Output file for the IDL (stdout if not specified)")
	idlCmd.AddCommand(idl_typeCmd)

	carapace.Gen(idl_typeCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
	})

	carapace.Gen(idl_typeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
