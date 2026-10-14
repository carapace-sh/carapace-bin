package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_shCmd = &cobra.Command{
	Use:   "sh",
	Short: "Run your shell in a `brew bundle exec` environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_shCmd).Standalone()

	bundle_shCmd.Flags().Bool("check", false, "Check that all dependencies in the Brewfile are installed before starting the shell. Enabled by default if `$HOMEBREW_BUNDLE_CHECK` is set.")
	bundle_shCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_shCmd.Flags().Bool("help", false, "Show this message.")
	bundle_shCmd.Flags().Bool("install", false, "Run `install` before starting the shell.")
	bundle_shCmd.Flags().Bool("no-secrets", false, "Attempt to remove secrets from the environment before starting the shell.")
	bundle_shCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_shCmd.Flags().Bool("services", false, "Temporarily start services while running the shell. Enabled by default if `$HOMEBREW_BUNDLE_SERVICES` is set.")
	bundle_shCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundleCmd.AddCommand(bundle_shCmd)
}
