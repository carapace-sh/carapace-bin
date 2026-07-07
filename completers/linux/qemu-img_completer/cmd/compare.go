package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Check if two images have the same contents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareCmd).Standalone()

	compareCmd.Flags().StringP("a-format", "f", "", "specify FILE1 image format explicitly")
	compareCmd.Flags().StringP("b-format", "F", "", "specify FILE2 image format explicitly")
	compareCmd.Flags().StringP("cache", "T", "", "images caching mode (default: writeback)")
	compareCmd.Flags().BoolP("force-share", "U", false, "open images in shared mode for concurrent access")
	compareCmd.Flags().Bool("image-opts", false, "treat FILE1 and FILE2 as option strings")
	compareCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	compareCmd.Flags().BoolP("progress", "p", false, "display progress information")
	compareCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	compareCmd.Flags().BoolP("strict", "s", false, "strict mode, also check if sizes are equal")
	rootCmd.AddCommand(compareCmd)

	carapace.Gen(compareCmd).FlagCompletion(carapace.ActionMap{
		"a-format": qemu.ActionImageFormats(),
		"b-format": qemu.ActionImageFormats(),
		"cache":    qemu.ActionCacheModes(),
	})

	carapace.Gen(compareCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
