package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fuzz_cminCmd = &cobra.Command{
	Use:   "cmin",
	Short: "Minimize corpus to smallest set preserving coverage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fuzz_cminCmd).Standalone()

	fuzz_cminCmd.Flags().String("corpus-in", "", "Input corpus directory (flag alternative)")
	fuzz_cminCmd.Flags().String("corpus-out", "", "Output directory (default: overwrite input)")
	fuzz_cminCmd.Flags().BoolP("help", "h", false, "Print help")
	fuzz_cminCmd.Flags().Bool("release", false, "Build in release mode")
	fuzzCmd.AddCommand(fuzz_cminCmd)

	carapace.Gen(fuzz_cminCmd).FlagCompletion(carapace.ActionMap{
		"corpus-in":  carapace.ActionFiles(),
		"corpus-out": carapace.ActionFiles(),
	})

	carapace.Gen(fuzz_cminCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
