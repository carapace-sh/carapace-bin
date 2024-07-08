package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/spf13/cobra"
)

var legacyCmd = &cobra.Command{
	Use:   "legacy",
	Aliases: []string{"l"},
	Short: "legacy entry-point command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func add_legacy_flags(cmd *cobra.Command) {
	cmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	cmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	cmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	cmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	cmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")

	cmd.Flags().Bool("help-ini", false, "show live configuration (default: False)")
	cmd.Flags().Bool("hi", false, "show live configuration (default: False)")
	cmd.MarkFlagsMutuallyExclusive("help-ini", "hi")

	cmd.Flags().Bool("showconfig", false, "show live configuration (by default all env, with -l only default targets, specific via TOXENV/-e) (default: False)")
	cmd.Flags().BoolP("listenvs-all", "a", false, "show list of all defined environments (with description if verbose) (default: False)")
	cmd.Flags().BoolP("listenvs", "l", false, "show list of test environments (with description if verbose) (default: False)")
	cmd.Flags().String("devenv", "", "sets up a development environment at ENVDIR based on the env's tox configuration specified by`-e` (-e defaults to py) (default: None)")
	cmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	cmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")

	// these two are the same flag; there are two long flags for this
	cmd.Flags().BoolP("pkg-only", "b", false, "only perform the packaging activity")
	cmd.Flags().Bool("sdistonly", false, "only perform the packaging activity")
	cmd.MarkFlagsMutuallyExclusive("pkg-only", "sdistonly")

	cmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	cmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	cmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	cmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")
	cmd.Flags().StringP("parallel", "p", "0", "run tox environments in parallel, the argument controls limit: all, auto - cpu count, some positive number, zero is turn off (default: 0)")
	cmd.Flags().BoolP("parallel-live", "o", false, "connect to stdout while running environments")
	cmd.Flags().Bool("parallel-no-spinner", false, "run tox environments in parallel, but don't show the spinner, implies --parallel")
	cmd.Flags().Bool("pre", false, "deprecated use PIP_PRE in set_env instead - install pre-releases and development versions ofdependencies; this will set PIP_PRE=1 environment variable (default: False)")
	cmd.Flags().StringArray("force-dep", []string{}, "Forces a certain version of one of the dependencies when configuring the virtual environment. REQ Examples 'pytest<6.1' or 'django>=2.2'. (default: [])")
	cmd.Flags().Bool("sitepackages", false, "deprecated use VIRTUALENV_SYSTEM_SITE_PACKAGES=1, override sitepackages setting to True in all envs (default: False)")
	cmd.Flags().Bool("alwayscopy", false, "deprecated use VIRTUALENV_ALWAYS_COPY=1, override always copy setting to True in all envs (default: False)")

	cmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	cmd.Flags().StringArrayS("m", "m", []string{}, "labels to evaluate (default: [])")
	cmd.Flags().StringArrayS("f", "f", []string{}, "factors to evaluate (passing multiple factors means 'AND', passing this option multiple times means 'OR') (default: [])")
	cmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"e":  tox.ActionEnvironments().UniqueList(","),
	})
}

func init() {
	carapace.Gen(legacyCmd).Standalone()
	
	add_common_flags(legacyCmd)
	add_legacy_flags(legacyCmd)

	rootCmd.AddCommand(legacyCmd)
}
