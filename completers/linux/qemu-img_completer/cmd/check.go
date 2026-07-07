package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check basic image integrity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().StringP("cache", "T", "", "cache mode (default: writeback)")
	checkCmd.Flags().BoolP("force-share", "U", false, "open image in shared mode for concurrent access")
	checkCmd.Flags().StringP("format", "f", "", "specifies the format of the image explicitly")
	checkCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	checkCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	checkCmd.Flags().String("output", "", "output format (default: human)")
	checkCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	checkCmd.Flags().StringP("repair", "r", "", "repair errors of the given category")
	rootCmd.AddCommand(checkCmd)

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"cache":  qemu.ActionCacheModes(),
		"format": qemu.ActionImageFormats(),
		"output": carapace.ActionValues("human", "json"),
		"repair": carapace.ActionValues("leaks", "all"),
	})

	carapace.Gen(checkCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
