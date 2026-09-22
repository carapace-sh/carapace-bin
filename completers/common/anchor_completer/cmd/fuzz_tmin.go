package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fuzz_tminCmd = &cobra.Command{
	Use:   "tmin",
	Short: "Minimize a crash to smallest reproducing action sequence",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fuzz_tminCmd).Standalone()

	fuzz_tminCmd.Flags().Bool("all", false, "Minimize all crashes for this test")
	fuzz_tminCmd.Flags().String("crash-meta-dir", "", "Custom directory for .meta.json files (default: same as crashes directory)")
	fuzz_tminCmd.Flags().BoolP("help", "h", false, "Print help")
	fuzz_tminCmd.Flags().Bool("release", false, "Build in release mode")
	fuzzCmd.AddCommand(fuzz_tminCmd)
}
