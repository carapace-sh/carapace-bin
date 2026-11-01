package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if all dependencies present in the `Brewfile` are installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_checkCmd).Standalone()

	bundle_checkCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_checkCmd.Flags().Bool("help", false, "Show this message.")
	bundle_checkCmd.Flags().Bool("install", false, "Run `install` before checking dependencies.")
	bundle_checkCmd.Flags().Bool("no-upgrade", false, "Do not check for outdated dependencies. Note they may still be upgraded by `brew install` if needed. Enabled by default if `$HOMEBREW_BUNDLE_NO_UPGRADE` is set.")
	bundle_checkCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_checkCmd.Flags().BoolP("verbose", "v", false, "List all missing dependencies.")
	bundleCmd.AddCommand(bundle_checkCmd)
}
