package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var randCmd = &cobra.Command{
	Use:     "rand",
	Short:   "Generate pseudo-random bytes",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(randCmd).Standalone()

	randCmd.Flags().BoolS("base64", "base64", false, "Base64 encode output")
	randCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	randCmd.Flags().BoolS("hex", "hex", false, "Hex encode output")
	randCmd.Flags().StringS("out", "out", "", "Output file")
	common.AddProviderFlags(randCmd)
	common.AddRandomStateFlags(randCmd)
	rootCmd.AddCommand(randCmd)

	carapace.Gen(randCmd).FlagCompletion(carapace.ActionMap{
		"engine": action.ActionEngines(),
		"out":    carapace.ActionFiles(),
	})
}
