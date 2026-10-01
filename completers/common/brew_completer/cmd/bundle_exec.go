package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var bundle_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run an external command in an isolated build environment based on the `Brewfile` dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_execCmd).Standalone()

	bundle_execCmd.Flags().SetInterspersed(false)

	bundle_execCmd.Flags().Bool("check", false, "Check that all dependencies in the Brewfile are installed before executing the command. Enabled by default if `$HOMEBREW_BUNDLE_CHECK` is set.")
	bundle_execCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bundle_execCmd.Flags().Bool("deny-network", false, "Deny network access from inside the sandbox.")
	bundle_execCmd.Flags().Bool("help", false, "Show this message.")
	bundle_execCmd.Flags().Bool("install", false, "Run `install` before executing the command.")
	bundle_execCmd.Flags().Bool("no-secrets", false, "Attempt to remove secrets from the environment before executing the command.")
	bundle_execCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bundle_execCmd.Flags().String("sandbox", "", "Run <command> in Homebrew's sandbox, allowing writes to <path> and Homebrew's temporary and cache directories.")
	bundle_execCmd.Flags().Bool("services", false, "Temporarily start services while executing the command. Enabled by default if `$HOMEBREW_BUNDLE_SERVICES` is set.")
	bundle_execCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bundleCmd.AddCommand(bundle_execCmd)

	carapace.Gen(bundle_execCmd).FlagCompletion(carapace.ActionMap{
		"sandbox": carapace.ActionDirectories(),
	})

	carapace.Gen(bundle_execCmd).PositionalCompletion(
		carapace.ActionExecutables(),
	)

	carapace.Gen(bundle_execCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
