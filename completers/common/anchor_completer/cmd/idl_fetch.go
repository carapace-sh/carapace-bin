package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches an IDL for the given program from a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_fetchCmd).Standalone()

	idl_fetchCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_fetchCmd.Flags().Bool("non-canonical", false, "Fetch non-canonical metadata account (third-party metadata)")
	idl_fetchCmd.Flags().StringP("out", "o", "", "Output file for the IDL (stdout if not specified)")
	idlCmd.AddCommand(idl_fetchCmd)

	carapace.Gen(idl_fetchCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
	})
}
