package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fuzz_showCmd = &cobra.Command{
	Use:   "show",
	Short: "View/replay crashes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fuzz_showCmd).Standalone()

	fuzz_showCmd.Flags().String("crash-meta-dir", "", "Custom directory for .meta.json files (default: same as crashes_dir)")
	fuzz_showCmd.Flags().String("crashes-dir", "", "Custom crashes directory to read from")
	fuzz_showCmd.Flags().BoolP("help", "h", false, "Print help")
	fuzz_showCmd.Flags().Bool("regen", false, "Batch-regenerate .meta.json for all crashes (requires --replay)")
	fuzz_showCmd.Flags().Bool("replay", false, "Actually replay the crash (requires compiled binary)")
	fuzzCmd.AddCommand(fuzz_showCmd)

	carapace.Gen(fuzz_showCmd).FlagCompletion(carapace.ActionMap{
		"crash-meta-dir": carapace.ActionDirectories(),
		"crashes-dir":    carapace.ActionDirectories(),
	})
}
