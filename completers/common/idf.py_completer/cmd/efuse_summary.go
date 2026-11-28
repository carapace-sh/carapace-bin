package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var efuseSummaryCmd = &cobra.Command{
	Use:   "efuse-summary",
	Short: "Show eFuse summary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseSummaryCmd).Standalone()

	common.AddEfuseFlags(efuseSummaryCmd)
	efuseSummaryCmd.Flags().String("format", "", "Summary format")

	rootCmd.AddCommand(efuseSummaryCmd)

	carapace.Gen(efuseSummaryCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "summary", "value-only"),
	})
}
