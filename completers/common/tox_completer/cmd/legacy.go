package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyCmd = &cobra.Command{
	Use:     "legacy",
	Aliases: []string{"l"},
	Short:   "legacy entry-point command",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func add_legacy_flags(cmd *cobra.Command) {
	cmd.Flags().String("devenv", "", "sets up a development environment at ENVDIR based on the env's tox configuration specified by`-e` (-e defaults to py) (default: None)")
	cmd.Flags().Bool("help-ini", false, "show live configuration (default: False)")
	cmd.Flags().Bool("hi", false, "show live configuration (default: False)")
	cmd.MarkFlagsMutuallyExclusive("help-ini", "hi")
	cmd.Flags().Bool("alwayscopy", false, "deprecated use VIRTUALENV_ALWAYS_COPY=1, override always copy setting to True in all envs (default: False)")
	cmd.Flags().StringArray("force-dep", nil, "Forces a certain version of one of the dependencies when configuring the virtual environment. REQ Examples 'pytest<6.1' or 'django>=2.2'. (default: [])")
	cmd.Flags().BoolP("listenvs", "l", false, "show list of test environments (with description if verbose) (default: False)")
	cmd.Flags().BoolP("listenvs-all", "a", false, "show list of all defined environments (with description if verbose) (default: False)")
	cmd.Flags().StringP("parallel", "p", "0", "run tox environments in parallel, the argument controls limit: all, auto - cpu count, some positive number, zero is turn off (default: 0)")
	cmd.Flags().BoolP("parallel-live", "o", false, "connect to stdout while running environments")
	cmd.Flags().Bool("parallel-no-spinner", false, "run tox environments in parallel, but don't show the spinner, implies --parallel")
	cmd.Flags().Bool("pre", false, "deprecated use PIP_PRE in set_env instead - install pre-releases and development versions ofdependencies; this will set PIP_PRE=1 environment variable (default: False)")
	cmd.Flags().Bool("showconfig", false, "show live configuration (by default all env, with -l only default targets, specific via TOXENV/-e) (default: False)")
	cmd.Flags().Bool("sitepackages", false, "deprecated use VIRTUALENV_SYSTEM_SITE_PACKAGES=1, override sitepackages setting to True in all envs (default: False)")
	addCommonSubcommandFlags(cmd)
	addPkgOnlyFlags(cmd)
	addEnvFilteringFlags(cmd)
	addEnvSelectFlag(cmd)
	addCommonRunFlags(cmd)
}

func init() {
	carapace.Gen(legacyCmd).Standalone()

	add_legacy_flags(legacyCmd)
	rootCmd.AddCommand(legacyCmd)
}
