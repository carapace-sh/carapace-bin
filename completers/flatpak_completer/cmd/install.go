package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install an application or runtime",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("app", false, "Look for app with the specified name")
	installCmd.Flags().String("arch", "", "Arch to install for")
	installCmd.Flags().BoolP("assumeyes", "y", false, "Automatically answer yes for all questions")
	installCmd.Flags().Bool("bundle", false, "Assume LOCATION is a .flatpak single-file bundle")
	installCmd.Flags().Bool("from", false, "Assume LOCATION is a .flatpakref application description")
	installCmd.Flags().String("gpg-file", "", "Check bundle signatures with GPG key from FILE (- for stdin)")
	installCmd.Flags().BoolP("help", "h", false, "Show help options")
	installCmd.Flags().Bool("include-debug", false, "Additionally install the debug info for the given refs and their dependencies")
	installCmd.Flags().Bool("include-sdk", false, "Additionally install the SDK used to build the given refs")
	installCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	installCmd.Flags().Bool("no-auto-pin", false, "Don't automatically pin explicit installs")
	installCmd.Flags().Bool("no-deploy", false, "Don't deploy, only download to local cache")
	installCmd.Flags().Bool("no-deps", false, "Don't verify/install runtime dependencies")
	installCmd.Flags().Bool("no-pull", false, "Don't pull, only install from local cache")
	installCmd.Flags().Bool("no-related", false, "Don't install related refs")
	installCmd.Flags().Bool("no-static-deltas", false, "Don't use static deltas")
	installCmd.Flags().Bool("noninteractive", false, "Produce minimal output and don't ask questions")
	installCmd.Flags().Bool("or-update", false, "Update install if already installed")
	installCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	installCmd.Flags().Bool("reinstall", false, "Uninstall first if already installed")
	installCmd.Flags().Bool("runtime", false, "Look for runtime with the specified name")
	installCmd.Flags().String("sideload-repo", "", "Use this local repo for sideloads")
	installCmd.Flags().String("subpath", "", "Only install this subpath")
	installCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	installCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	installCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"arch":     flatpak.ActionArches(),
		"gpg-file": carapace.ActionFiles(),
		// "installation":  carapace.ActionValues(), TODO
		"sideload-repo": carapace.ActionFiles(),
		// "subpath":       carapace.ActionValues(), TODO
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		flatpak.ActionApplicationSearch().FilterArgs(),
	)
}
