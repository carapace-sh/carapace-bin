package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize the image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resizeCmd).Standalone()

	resizeCmd.Flags().StringP("format", "f", "", "specify FILE format explicitly")
	resizeCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	resizeCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	resizeCmd.Flags().String("preallocation", "", "specify FMT-specific preallocation type for the new areas")
	resizeCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	resizeCmd.Flags().Bool("shrink", false, "allow operation when the new size is smaller than the original")
	rootCmd.AddCommand(resizeCmd)

	carapace.Gen(resizeCmd).FlagCompletion(carapace.ActionMap{
		"format":        qemu.ActionImageFormats(),
		"preallocation": carapace.ActionValues("off", "falloc", "full"),
	})

	carapace.Gen(resizeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
