package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyCmd = &cobra.Command{
	Use:   "tox",
	Short: "legacy entry-point command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyCmd).Standalone()

	// TODO: process this list
	legacyCmd.Flags().StringS("e", "e", "", "enumerate (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	legacyCmd.Flags().Bool("alwayscopy", false, "deprecated use VIRTUALENV_ALWAYS_COPY=1, override always copy setting to True in all envs (default: False)")
	legacyCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	legacyCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	legacyCmd.Flags().String("devenv", "", "sets up a development environment at ENVDIR based on the env's tox configuration specified by`-e` (-e defaults to py) (default: None)")
	legacyCmd.Flags().String("exit-and-dump-after", "", "")
	legacyCmd.Flags().String("force-dep", "", "Forces a certain version of one of the dependencies when configuring the virtual environment. REQ Examples 'pytest<6.1' or 'django>=2.2'. (default: [])")
	legacyCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 2745445130)")
	legacyCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	legacyCmd.Flags().String("help-ini,", "", "show live configuration (default: False)")
	legacyCmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	legacyCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	legacyCmd.Flags().BoolP("listenvs", "l", false, "show list of test environments (with description if verbose) (default: False)")
	legacyCmd.Flags().BoolP("listenvs-all", "a", false, "show list of all defined environments (with description if verbose) (default: False)")
	legacyCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	legacyCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	legacyCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	legacyCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	legacyCmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")
	legacyCmd.Flags().BoolP("parallel-live", "o", false, "connect to stdout while running environments (default: False)")
	legacyCmd.Flags().Bool("parallel-no-spinner", false, "run tox environments in parallel, but don't show the spinner, implies --parallel (default: False)")
	legacyCmd.Flags().StringP("pkg-only,", "b", "", "")
	legacyCmd.Flags().Bool("pre", false, "deprecated use PIP_PRE in set_env instead - install pre-releases and development versions ofdependencies; this will set PIP_PRE=1 environment variable (default: False)")
	legacyCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	legacyCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	legacyCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	legacyCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	legacyCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	legacyCmd.Flags().Bool("showconfig", false, "show live configuration (by default all env, with -l only default targets, specific via TOXENV/-e) (default: False)")
	legacyCmd.Flags().Bool("sitepackages", false, "deprecated use VIRTUALENV_SYSTEM_SITE_PACKAGES=1, override sitepackages setting to True in all envs (default: False)")
	legacyCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: )")
	legacyCmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")
	legacyCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	legacyCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	legacyCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(legacyCmd)
}
