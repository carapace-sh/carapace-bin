package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/spf13/cobra"
)

// add_common_flags defines the flags on a tox subcommand that are common to all.
func add_common_flags(cmd *cobra.Command) {
	cmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	cmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	cmd.Flags().Uint("exit-and-dump-after", 0, "dump tox threads after n seconds and exit the app - useful to debug when tox hangs, 0 means disabled (default: 0)")
	cmd.Flags().StringP("conf", "c", "", "configuration file/folder for tox (if not specified will discover one) (default: None)")
	cmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")
	cmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	cmd.Flags().String("runner", "virtualenv", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	cmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	cmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	cmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	cmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	cmd.Flags().StringArrayP("override", "x", []string{}, "configuration override(s), e.g., -x testenv:pypy3.ignore_errors=True (default: [])")

	cmd.Flags().BoolSliceP("quiet", "q", []bool{}, "decrease verbosity (default: 0)")
	cmd.Flags().BoolSliceP("verbose", "v", []bool{}, "increase verbosity (default: 2)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"colored":  carapace.ActionValues("yes", "no"),
	})
}

// add_common_subcommand_flags defines the flags on a tox subcommand that are common across a majority of subcommands.
func add_common_subcommand_flags(cmd *cobra.Command) {
	cmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	cmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	cmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	cmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	cmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
}

// add_pkg_only_flags adds the pkg-only / sdistonly group of flags
func add_pkg_only_flags(cmd *cobra.Command) {
	// these two are the same flag; there are two long flags for this
	cmd.Flags().BoolP("pkg-only", "b", false, "only perform the packaging activity")
	cmd.Flags().Bool("sdistonly", false, "only perform the packaging activity")
	cmd.MarkFlagsMutuallyExclusive("pkg-only", "sdistonly")
}

// add_env_filtering_flags adds the -m, -f, and --skip-env flags
func add_env_filtering_flags(cmd *cobra.Command) {
	cmd.Flags().StringArrayS("m", "m", []string{}, "labels to evaluate (default: [])")
	cmd.Flags().StringArrayS("f", "f", []string{}, "factors to evaluate (passing multiple factors means 'AND', passing this option multiple times means 'OR') (default: [])")
	cmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")
}

// add_env_select_flag adds the -e flag for multiple environments
func add_env_select_flag(cmd *cobra.Command) {
	cmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"e":  tox.ActionEnvironments().UniqueList(","),
	})
}

// add_common_run_flags adds flags common to subcommonds that run something in an environment
func add_common_run_flags(cmd *cobra.Command) {
	cmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	cmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")
	cmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	cmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	cmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	cmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"skip-missing-interpreters":  carapace.ActionValues("config", "true", "false"),
	})
}
