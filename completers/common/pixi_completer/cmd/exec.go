package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Run a command and install it in a temporary environment",
	Aliases: []string{"x"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	execCmd.Flags().StringSliceP("channel", "c", nil, "The channels to consider as a name or a url. Multiple channels can be specified by using this field multiple times")
	execCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	execCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	execCmd.Flags().Bool("force-reinstall", false, "If specified a new environment is always created even if one already exists")
	execCmd.Flags().String("list", "", "Before executing the command, list packages in the environment Specify `--list=some_regex` to filter the shown packages")
	execCmd.Flags().Bool("no-modify-ps1", false, "Disable modification of the PS1 prompt to indicate the temporary environment")
	execCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	execCmd.Flags().StringP("platform", "p", "", "The platform to create the environment for")
	execCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	execCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	execCmd.Flags().StringSliceP("spec", "s", nil, "Matchspecs of package to install. If this is not provided, the package is guessed from the command")
	execCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	execCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	execCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	execCmd.Flags().StringSliceP("with", "w", nil, "Matchspecs of package to install, while also guessing a package from the command")
	execCmd.Flag("list").NoOptDefVal = " "
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
