package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "tox",
	Short: "automation project",
	Long:  "https://tox.wiki/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	rootCmd.PersistentFlags().StringP("conf", "c", "", "configuration file/folder for tox (if not specified will discover one) (default: None)")
	rootCmd.PersistentFlags().Uint("exit-and-dump-after", 0, "dump tox threads after n seconds and exit the app - useful to debug when tox hangs, 0 means disabled (default: 0)")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.PersistentFlags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	rootCmd.PersistentFlags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	rootCmd.PersistentFlags().StringArrayP("override", "x", nil, "configuration override(s), e.g., -x testenv:pypy3.ignore_errors=True (default: [])")
	rootCmd.PersistentFlags().CountP("quiet", "q", "decrease verbosity (default: 0)")
	rootCmd.PersistentFlags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	rootCmd.PersistentFlags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	rootCmd.PersistentFlags().String("runner", "virtualenv", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	rootCmd.PersistentFlags().Bool("version", false, "show program's and plugins version number and exit")
	rootCmd.PersistentFlags().CountP("verbose", "v", "increase verbosity (default: 2)")
	rootCmd.PersistentFlags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")
	add_legacy_flags(rootCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"colored": carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"conf":    carapace.ActionFiles(),
		"root":    carapace.ActionDirectories(),
		"workdir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("workdir")))
	})
}

// addCommonSubcommandFlags defines the flags on a tox subcommand that are common across a majority of subcommands.
func addCommonSubcommandFlags(cmd *cobra.Command) {
	cmd.Flags().StringArray("discover", nil, "for Python discovery first try these Python executables (default: [])")
	cmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	cmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	cmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	cmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")

	cmd.Flag("discover").Nargs = -1

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"discover":    carapace.ActionFiles(),
		"result-json": carapace.ActionFiles(),
	})
}

// addPkgOnlyFlags adds the pkg-only / sdistonly group of flags
func addPkgOnlyFlags(cmd *cobra.Command) {
	// these two are the same flag; there are two long flags for this
	cmd.Flags().BoolP("pkg-only", "b", false, "only perform the packaging activity")
	cmd.Flags().Bool("sdistonly", false, "only perform the packaging activity")
	cmd.MarkFlagsMutuallyExclusive("pkg-only", "sdistonly")
}

// addEnvFilteringFlags adds the -m, -f, and --skip-env flags
func addEnvFilteringFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayS("f", "f", nil, "factors to evaluate (passing multiple factors means 'AND', passing this option multiple times means 'OR') (default: [])")
	cmd.Flags().StringArrayS("m", "m", nil, "labels to evaluate (default: [])")
	cmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")

	cmd.Flag("f").Nargs = -1
	cmd.Flag("m").Nargs = -1
}

// addEnvSelectFlag adds the -e flag for multiple environments
func addEnvSelectFlag(cmd *cobra.Command) {
	cmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"e": tox.ActionEnvironments().UniqueList(","),
	})
}

// addCommonRunFlags adds flags common to subcommonds that run something in an environment
func addCommonRunFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	cmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	cmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	cmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")
	cmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	cmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"skip-missing-interpreters": carapace.ActionValues("config", "true", "false").StyleF(style.ForKeyword),
	})
}
