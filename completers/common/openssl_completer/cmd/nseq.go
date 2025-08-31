package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var nseqCmd = &cobra.Command{
	Use:     "nseq",
	Short:   "Create or examine a Netscape certificate sequence",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nseqCmd).Standalone()

	nseqCmd.Flags().StringS("in", "in", "", "Input file")
	nseqCmd.Flags().StringS("out", "out", "", "Output file")
	nseqCmd.Flags().BoolS("toseq", "toseq", false, "Output NS Sequence file")
	common.AddProviderFlags(nseqCmd)
	rootCmd.AddCommand(nseqCmd)

	carapace.Gen(nseqCmd).FlagCompletion(carapace.ActionMap{
		"in":  carapace.ActionFiles(),
		"out": carapace.ActionFiles(),
	})
}
