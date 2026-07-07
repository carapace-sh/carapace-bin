package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Change the backing file of the image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().StringP("backing", "b", "", "rebase onto this file")
	rebaseCmd.Flags().StringP("backing-cache", "T", "", "BACKING_FILE cache mode (default: writeback)")
	rebaseCmd.Flags().StringP("backing-format", "B", "", "specify format for BACKING_FILE explicitly")
	rebaseCmd.Flags().BoolP("backing-unsafe", "u", false, "do not fail if BACKING_FILE can not be read")
	rebaseCmd.Flags().StringP("cache", "t", "", "cache mode for FILE (default: writeback)")
	rebaseCmd.Flags().BoolP("compress", "c", false, "compress image")
	rebaseCmd.Flags().BoolP("force-share", "U", false, "open image in shared mode for concurrent access")
	rebaseCmd.Flags().StringP("format", "f", "", "specify FILE format explicitly")
	rebaseCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	rebaseCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	rebaseCmd.Flags().BoolP("progress", "p", false, "display progress information")
	rebaseCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.AddCommand(rebaseCmd)

	carapace.Gen(rebaseCmd).FlagCompletion(carapace.ActionMap{
		"backing":        carapace.ActionFiles(),
		"backing-cache":  qemu.ActionCacheModes(),
		"backing-format": qemu.ActionImageFormats(),
		"cache":          qemu.ActionCacheModes(),
		"format":         qemu.ActionImageFormats(),
	})

	carapace.Gen(rebaseCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
