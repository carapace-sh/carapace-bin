package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hdiutil",
	Short: "manipulate disk images",
	Long:  "https://keith.github.io/xcode-manpages/hdiutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"attach", "Attach a disk image",
			"detach", "Detach a disk image",
			"create", "Create a disk image",
			"convert", "Convert a disk image from one format to another",
			"burn", "Burn a disk image to disc",
			"verify", "Verify the checksum of a disk image",
			"checksum", "Compute the checksum of a disk image",
			"imageinfo", "Display information about a disk image",
			"info", "Display information about attached disk images",
			"compact", "Compact a sparse disk image",
			"resize", "Resize a disk image",
			"chpass", "Change the password of an encrypted disk image",
			"makehybrid", "Create a hybrid disk image",
			"mount", "Mount a disk image",
			"unmount", "Unmount a disk image",
			"eject", "Eject a disk image",
			"help", "Display help",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
