package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var bundle_installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"upgrade"},
	Short:   "Install and upgrade (by default) all dependencies from the `Brewfile`",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_installCmd).Standalone()

	bundle_installCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_installCmd.Flags().BoolP("force", "f", false, "Run with `--force`/`--overwrite`.")
	bundle_installCmd.Flags().Bool("force-cleanup", false, "Perform cleanup after installing dependencies without asking. Enabled by default if `$HOMEBREW_BUNDLE_FORCE_INSTALL_CLEANUP` is set and `--global` is passed.")
	bundle_installCmd.Flags().Bool("help", false, "Show this message.")
	bundle_installCmd.Flags().String("jobs", "", "Run up to this many formula installations in parallel. Defaults to 1 (sequential). Use `auto` for the number of CPU cores (max 4).")
	bundle_installCmd.Flags().Bool("no-upgrade", false, "Do not run `brew upgrade` on outdated dependencies. Note they may still be upgraded by `brew install` if needed. Enabled by default if `$HOMEBREW_BUNDLE_NO_UPGRADE` is set.")
	bundle_installCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_installCmd.Flags().Bool("upgrade", false, "Run `brew upgrade` on outdated dependencies, even if `$HOMEBREW_BUNDLE_NO_UPGRADE` is set.")
	bundle_installCmd.Flags().String("upgrade-formulae", "", "Run `brew upgrade` on any of these comma-separated formulae, even if `$HOMEBREW_BUNDLE_NO_UPGRADE` is set.")
	bundle_installCmd.Flags().BoolP("verbose", "v", false, "Print output from commands as they are run.")
	bundle_installCmd.Flags().Bool("zap", false, "Use `zap` instead of `uninstall` when cleaning up casks after installing dependencies.")
	bundleCmd.AddCommand(bundle_installCmd)

	carapace.Gen(bundle_installCmd).FlagCompletion(carapace.ActionMap{
		"jobs":             carapace.ActionValues("auto"),
		"upgrade-formulae": brew.ActionAllFormulae().UniqueList(","),
	})
}
