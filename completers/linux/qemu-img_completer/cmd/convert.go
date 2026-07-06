package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/qemu-img_completer/cmd/action"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Copy one or more images to another with optional format conversion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()

	convertCmd.Flags().StringP("backing", "b", "", "create target image to be a CoW on top of BACKING_FILE")
	convertCmd.Flags().StringP("backing-format", "F", "", "specifies the format of BACKING_FILE")
	convertCmd.Flags().Bool("bitmaps", false, "also copy any persistent bitmaps present in source")
	convertCmd.Flags().BoolP("compress", "c", false, "create compressed output image (qcow and qcow2 formats only)")
	convertCmd.Flags().BoolP("copy-range-offloading", "C", false, "try to use copy offloading")
	convertCmd.Flags().BoolP("force-share", "U", false, "open images in shared mode for concurrent access")
	convertCmd.Flags().StringP("format", "f", "", "specify format of all SRC_FILEs explicitly")
	convertCmd.Flags().Bool("image-opts", false, "treat each SRC_FILE as an option string")
	convertCmd.Flags().BoolP("no-create", "n", false, "omit target volume creation")
	convertCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	convertCmd.Flags().BoolP("oob-writes", "W", false, "enable out-of-order writes to improve performance")
	convertCmd.Flags().StringP("parallel", "m", "", "specify parallelism (default: 8)")
	convertCmd.Flags().BoolP("progress", "p", false, "display progress information")
	convertCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	convertCmd.Flags().StringP("rate-limit", "r", "", "I/O rate limit, in bytes per second")
	convertCmd.Flags().Bool("salvage", false, "ignore errors on input")
	convertCmd.Flags().Bool("skip-broken-bitmaps", false, "skip any broken bitmaps")
	convertCmd.Flags().StringP("snapshot", "l", "", "specify source snapshot")
	convertCmd.Flags().StringP("source-cache", "T", "", "source image(s) cache mode")
	convertCmd.Flags().StringP("sparse-size", "S", "", "sparse size")
	convertCmd.Flags().StringP("target-cache", "t", "", "cache mode when opening output image")
	convertCmd.Flags().StringP("target-format", "O", "", "specify TGT_FILE image format (default: raw)")
	convertCmd.Flags().StringP("target-format-options", "o", "", "TGT_FMT-specific options")
	convertCmd.Flags().Bool("target-image-opts", false, "treat TGT_FILE as an option string")
	convertCmd.Flags().Bool("target-is-zero", false, "indicates that the target volume is pre-zeroed")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"backing":               carapace.ActionFiles(),
		"backing-format":        action.ActionImageFormats(),
		"format":                action.ActionImageFormats(),
		"source-cache":          action.ActionCacheModes(),
		"target-cache":          action.ActionCacheModes(),
		"target-format":         action.ActionImageFormats(),
		"target-format-options": carapace.ActionValues(),
	})

	carapace.Gen(convertCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
