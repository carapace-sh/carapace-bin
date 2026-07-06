package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/qemu-img_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "List or manipulate snapshots in the image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	snapshotCmd.Flags().StringP("apply", "a", "", "apply named snapshot to the base")
	snapshotCmd.Flags().StringP("create", "c", "", "create named snapshot")
	snapshotCmd.Flags().StringP("delete", "d", "", "delete named snapshot")
	snapshotCmd.Flags().BoolP("force-share", "U", false, "open image in shared mode for concurrent access")
	snapshotCmd.Flags().StringP("format", "f", "", "specify FILE format explicitly")
	snapshotCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	snapshotCmd.Flags().BoolP("list", "l", false, "list snapshots in FILE")
	snapshotCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	snapshotCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.AddCommand(snapshotCmd)

	carapace.Gen(snapshotCmd).FlagCompletion(carapace.ActionMap{
		"format": action.ActionImageFormats(),
	})

	carapace.Gen(snapshotCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
