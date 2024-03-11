package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var remoteInfoCmd = &cobra.Command{
	Use:     "remote-info [OPTION…]  REMOTE REF",
	Short:   "Show information about an application or runtime in a remote",
	GroupID: "repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteInfoCmd).Standalone()

	remoteInfoCmd.Flags().Bool("app", false, "Look for app with the specified name")
	remoteInfoCmd.Flags().String("arch", "", "Arch to install for")
	remoteInfoCmd.Flags().Bool("cached", false, "Use local caches even if they are stale")
	remoteInfoCmd.Flags().String("commit", "", "Commit to show info for")
	remoteInfoCmd.Flags().BoolP("help", "h", false, "Show help options")
	remoteInfoCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	remoteInfoCmd.Flags().Bool("log", false, "Display log")
	remoteInfoCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	remoteInfoCmd.Flags().Bool("runtime", false, "Look for runtime with the specified name")
	remoteInfoCmd.Flags().BoolP("show-commit", "c", false, "Show commit")
	remoteInfoCmd.Flags().BoolP("show-metadata", "m", false, "Show metadata")
	remoteInfoCmd.Flags().BoolP("show-parent", "p", false, "Show parent")
	remoteInfoCmd.Flags().BoolP("show-ref", "r", false, "Show ref")
	remoteInfoCmd.Flags().Bool("show-runtime", false, "Show runtime")
	remoteInfoCmd.Flags().Bool("show-sdk", false, "Show sdk")
	remoteInfoCmd.Flags().Bool("sideloaded", false, "Only list refs available as sideloads")
	remoteInfoCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	remoteInfoCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	remoteInfoCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(remoteInfoCmd)

	/// TODO flag
	carapace.Gen(remoteInfoCmd).FlagCompletion(carapace.ActionMap{
		"arch": flatpak.ActionArches(),
		// "commit":       carapace.ActionValues(),
		// "installation": carapace.ActionValues(),
	})

	// TODO positional
}
