package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update [OPTION…] [REF…]",
	Short:   "Update applications or runtimes",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("app", false, "Look for app with the specified name")
	updateCmd.Flags().Bool("appstream", false, "Update appstream for remote")
	updateCmd.Flags().String("arch", "", "Arch to update for")
	updateCmd.Flags().BoolP("assumeyes", "y", false, "Automatically answer yes for all questions")
	updateCmd.Flags().String("commit", "", "Commit to deploy")
	updateCmd.Flags().Bool("force-remove", false, "Remove old files even if running")
	updateCmd.Flags().BoolP("help", "h", false, "Show help options")
	updateCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	updateCmd.Flags().Bool("no-deploy", false, "Don't deploy, only download to local cache")
	updateCmd.Flags().Bool("no-deps", false, "Don't verify/install runtime dependencies")
	updateCmd.Flags().Bool("no-pull", false, "Don't pull, only update from local cache")
	updateCmd.Flags().Bool("no-related", false, "Don't update related refs")
	updateCmd.Flags().Bool("no-static-deltas", false, "Don't use static deltas")
	updateCmd.Flags().Bool("noninteractive", false, "Produce minimal output and don't ask questions")
	updateCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	updateCmd.Flags().Bool("runtime", false, "Look for runtime with the specified name")
	updateCmd.Flags().String("sideload-repo", "", "Use this local repo for sideloads")
	updateCmd.Flags().String("subpath", "", "Only update this subpath")
	updateCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	updateCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	updateCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(updateCmd)

	// TODO flags
	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"arch": flatpak.ActionArches(),
		// "commit":        carapace.ActionValues(),
		// "installation":  carapace.ActionValues(),
		// "sideload-repo": carapace.ActionValues(),
		// "subpath":       carapace.ActionValues(),
	})

	// TODO positional
}
