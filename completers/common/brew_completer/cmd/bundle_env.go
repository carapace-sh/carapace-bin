package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print the environment variables that would be set in a `brew bundle exec` environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_envCmd).Standalone()

	bundle_envCmd.Flags().Bool("check", false, "Check that all dependencies in the Brewfile are installed before printing the environment. Enabled by default if `$HOMEBREW_BUNDLE_CHECK` is set.")
	bundle_envCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_envCmd.Flags().Bool("help", false, "Show this message.")
	bundle_envCmd.Flags().Bool("install", false, "Run `install` before printing the environment.")
	bundle_envCmd.Flags().Bool("no-secrets", false, "Attempt to remove secrets from the environment before printing it.")
	bundle_envCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_envCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundleCmd.AddCommand(bundle_envCmd)
}
