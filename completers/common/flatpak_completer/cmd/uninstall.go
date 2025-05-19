package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall [OPTION…] [REF…]",
	Short:   "Uninstall applications or runtimes",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().Bool("all", false, "Uninstall all")
	uninstallCmd.Flags().Bool("app", false, "Look for app with the specified name")
	uninstallCmd.Flags().String("arch", "", "Arch to uninstall")
	uninstallCmd.Flags().BoolP("assumeyes", "y", false, "Automatically answer yes for all questions")
	uninstallCmd.Flags().Bool("delete-data", false, "Delete app data")
	uninstallCmd.Flags().Bool("force-remove", false, "Remove files even if running")
	uninstallCmd.Flags().BoolP("help", "h", false, "Show help options")
	uninstallCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	uninstallCmd.Flags().Bool("keep-ref", false, "Keep ref in local repository")
	uninstallCmd.Flags().Bool("no-related", false, "Don't uninstall related refs")
	uninstallCmd.Flags().Bool("noninteractive", false, "Produce minimal output and don't ask questions")
	uninstallCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	uninstallCmd.Flags().Bool("runtime", false, "Look for runtime with the specified name")
	uninstallCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	uninstallCmd.Flags().Bool("unused", false, "Uninstall unused")
	uninstallCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	uninstallCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(uninstallCmd)

	// TODO flags
	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"arch": flatpak.ActionArches(),
		// "installation": carapace.ActionValues(),
	})

	// TODO positional
}
