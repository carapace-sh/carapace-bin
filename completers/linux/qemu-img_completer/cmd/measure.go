package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/qemu-img_completer/cmd/action"
	"github.com/spf13/cobra"
)

var measureCmd = &cobra.Command{
	Use:   "measure",
	Short: "Calculate the file size required for a new image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(measureCmd).Standalone()

	measureCmd.Flags().BoolP("force-share", "U", false, "open images in shared mode for concurrent access")
	measureCmd.Flags().StringP("format", "f", "", "specify format of FILE explicitly")
	measureCmd.Flags().Bool("image-opts", false, "indicates that FILE is a complete image specification")
	measureCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	measureCmd.Flags().String("output", "", "output format (default: human)")
	measureCmd.Flags().StringP("size", "s", "", "measure file size for given image size")
	measureCmd.Flags().StringP("snapshot", "l", "", "use this snapshot in FILE as source")
	measureCmd.Flags().StringP("target-format", "O", "", "desired target/output image format (default: raw)")
	measureCmd.Flags().StringP("target-format-options", "o", "", "options specific to TARGET_FMT")
	rootCmd.AddCommand(measureCmd)

	carapace.Gen(measureCmd).FlagCompletion(carapace.ActionMap{
		"format":                action.ActionImageFormats(),
		"output":                carapace.ActionValues("human", "json"),
		"target-format":         action.ActionImageFormats(),
		"target-format-options": carapace.ActionValues(),
	})

	carapace.Gen(measureCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
