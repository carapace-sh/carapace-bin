package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create and format a new image file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringP("backing", "b", "", "create target image to be a CoW on top of BACKING_FILE")
	createCmd.Flags().StringP("backing-format", "B", "", "specifies the format of BACKING_FILE")
	createCmd.Flags().BoolP("backing-unsafe", "u", false, "do not fail if BACKING_FILE can not be read")
	createCmd.Flags().StringP("format", "f", "", "specifies the format of the new image (default: raw)")
	createCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	createCmd.Flags().StringP("options", "o", "", "format-specific options")
	createCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"backing":        carapace.ActionFiles(),
		"backing-format": qemu.ActionImageFormats(),
		"format":         qemu.ActionImageFormats(),
	})

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
