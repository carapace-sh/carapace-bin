package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "[DEPRECATED] Fetch the IDL from a legacy on-chain IdlAccount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_fetchCmd).Standalone()

	legacyIdl_fetchCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_fetchCmd.Flags().StringP("out", "o", "", "Path to write the fetched IDL. Prints to stdout if omitted")
	legacyIdlCmd.AddCommand(legacyIdl_fetchCmd)

	carapace.Gen(legacyIdl_fetchCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
	})
}
