package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var ddCmd = &cobra.Command{
	Use:   "dd",
	Short: "Copy input to output with optional format conversion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ddCmd).Standalone()

	ddCmd.Flags().BoolP("force-share", "U", false, "open images in shared mode for concurrent access")
	ddCmd.Flags().StringP("format", "f", "", "specify format for INPUT explicitly")
	ddCmd.Flags().Bool("image-opts", false, "treat INPUT as an option string")
	ddCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	ddCmd.Flags().StringP("output-format", "O", "", "format of the OUTPUT (default: raw)")
	rootCmd.AddCommand(ddCmd)

	carapace.Gen(ddCmd).FlagCompletion(carapace.ActionMap{
		"format":        qemu.ActionImageFormats(),
		"output-format": qemu.ActionImageFormats(),
	})

	carapace.Gen(ddCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("bs=", "count=", "if=", "of=").NoSpace()
		}),
	)
}
