package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var bitmapCmd = &cobra.Command{
	Use:   "bitmap",
	Short: "Perform modifications of the persistent bitmap in the image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bitmapCmd).Standalone()

	bitmapCmd.Flags().Bool("add", false, "creates BITMAP in FILE")
	bitmapCmd.Flags().Bool("clear", false, "clears BITMAP in FILE")
	bitmapCmd.Flags().Bool("disable", false, "stops recording future edits to BITMAP in FILE")
	bitmapCmd.Flags().Bool("enable", false, "starts recording future edits to BITMAP in FILE")
	bitmapCmd.Flags().StringP("format", "f", "", "specify FILE format explicitly")
	bitmapCmd.Flags().StringP("granularity", "g", "", "sets non-default granularity for the bitmap being added")
	bitmapCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	bitmapCmd.Flags().String("merge", "", "merges contents of the SOURCE bitmap into BITMAP in FILE")
	bitmapCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	bitmapCmd.Flags().Bool("remove", false, "removes BITMAP from FILE")
	bitmapCmd.Flags().StringP("source-file", "b", "", "select alternative source file for --merge")
	bitmapCmd.Flags().StringP("source-format", "F", "", "specify format for SRC_FILE explicitly")
	rootCmd.AddCommand(bitmapCmd)

	carapace.Gen(bitmapCmd).FlagCompletion(carapace.ActionMap{
		"format":        qemu.ActionImageFormats(),
		"source-file":   carapace.ActionFiles(),
		"source-format": qemu.ActionImageFormats(),
	})

	carapace.Gen(bitmapCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
