package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command and install it in a temporary environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	execCmd.Flags().StringP("channel", "c", "", "The channels to consider as a name or a url")
	execCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	execCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	execCmd.Flags().Bool("force-reinstall", false, "Always create a new environment even if one already exists")
	execCmd.Flags().String("list", "", "List packages in the environment")
	execCmd.Flags().Bool("no-modify-ps1", false, "Disable modification of the PS1 prompt")
	execCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	execCmd.Flags().StringP("platform", "p", "", "The platform to create the environment for")
	execCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	execCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	execCmd.Flags().StringP("spec", "s", "", "Matchspecs of package to install")
	execCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	execCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	execCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	execCmd.Flags().StringP("with", "w", "", "Matchspecs of package to install, while also guessing a package from the command")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
