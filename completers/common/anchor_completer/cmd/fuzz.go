package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "Coverage-guided fuzzing for Solana programs (powered by Crucible)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fuzzCmd).Standalone()

	fuzzCmd.PersistentFlags().StringP("harness-dir", "C", "", "Use this directory instead of ./fuzz/<program_name>/")
	fuzzCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(fuzzCmd)

	carapace.Gen(fuzzCmd).FlagCompletion(carapace.ActionMap{
		"harness-dir": carapace.ActionDirectories(),
	})
}
