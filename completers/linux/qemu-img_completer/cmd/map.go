package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Dump image metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mapCmd).Standalone()

	mapCmd.Flags().BoolP("force-share", "U", false, "open image in shared mode for concurrent access")
	mapCmd.Flags().StringP("format", "f", "", "specify FILE image format explicitly")
	mapCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	mapCmd.Flags().String("max-length", "", "process at most LENGTH bytes")
	mapCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	mapCmd.Flags().String("output", "", "specify output format name (default: human)")
	mapCmd.Flags().StringP("start-offset", "s", "", "start at the given OFFSET in the image")
	rootCmd.AddCommand(mapCmd)

	carapace.Gen(mapCmd).FlagCompletion(carapace.ActionMap{
		"format": qemu.ActionImageFormats(),
		"output": carapace.ActionValues("human", "json"),
	})

	carapace.Gen(mapCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
