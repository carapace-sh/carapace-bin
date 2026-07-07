package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about the image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("backing-chain", false, "display information about the backing chain")
	infoCmd.Flags().StringP("cache", "t", "", "cache mode for FILE (default: writeback)")
	infoCmd.Flags().BoolP("force-share", "U", false, "open image in shared mode for concurrent access")
	infoCmd.Flags().StringP("format", "f", "", "specify FILE image format explicitly")
	infoCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	infoCmd.Flags().Bool("limits", false, "show detected block limits")
	infoCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	infoCmd.Flags().String("output", "", "specify output format (default: human)")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"cache":  qemu.ActionCacheModes(),
		"format": qemu.ActionImageFormats(),
		"output": carapace.ActionValues("human", "json"),
	})

	carapace.Gen(infoCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
