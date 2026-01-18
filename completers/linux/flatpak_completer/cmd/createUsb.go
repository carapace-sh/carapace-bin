package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var createUsbCmd = &cobra.Command{
	Use:     "create-usb [OPTION…] MOUNT-PATH [REF…]",
	Short:   "Copy apps or runtimes onto removable media",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createUsbCmd).Standalone()

	createUsbCmd.Flags().Bool("allow-partial", false, "Allow partial commits in the created repo")
	createUsbCmd.Flags().Bool("app", false, "Look for app with the specified name")
	createUsbCmd.Flags().String("arch", "", "Arch to copy")
	createUsbCmd.Flags().String("destination-repo", "", "Use custom repository directory within the mount")
	createUsbCmd.Flags().BoolP("help", "h", false, "Show help options")
	createUsbCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	createUsbCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	createUsbCmd.Flags().Bool("runtime", false, "Look for runtime with the specified name")
	createUsbCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	createUsbCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	createUsbCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(createUsbCmd)

	// TODO flags
	carapace.Gen(createUsbCmd).FlagCompletion(carapace.ActionMap{
		"arch": flatpak.ActionArches(),
		// "destination-repo": carapace.ActionValues(),
		// "installation": carapace.ActionValues(),
	})

	carapace.Gen(createUsbCmd).PositionalCompletion(
		carapace.Batch(
			fs.ActionMounts(),
			carapace.ActionFiles(),
		).ToA(),
		// TODO ref
	)
}
