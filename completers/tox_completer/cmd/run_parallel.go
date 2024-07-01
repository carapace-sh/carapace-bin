package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runParallelCmd = &cobra.Command{
	Use:   "run-parallel",
	Short: "run environments in parallel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runParallelCmd).Standalone()

	runParallelCmd.Flags().StringS("e", "e", "", "enumerate (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	runParallelCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	runParallelCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	runParallelCmd.Flags().String("exit-and-dump-after", "", "")
	runParallelCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 621153845)")
	runParallelCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	runParallelCmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	runParallelCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	runParallelCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	runParallelCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	runParallelCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	runParallelCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	runParallelCmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")
	runParallelCmd.Flags().BoolP("parallel-live", "o", false, "connect to stdout while running environments (default: False)")
	runParallelCmd.Flags().Bool("parallel-no-spinner", false, "run tox environments in parallel, but don't show the spinner, implies --parallel (default: False)")
	runParallelCmd.Flags().StringP("pkg-only,", "b", "", "")
	runParallelCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	runParallelCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	runParallelCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	runParallelCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	runParallelCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	runParallelCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: )")
	runParallelCmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")
	runParallelCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	runParallelCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	runParallelCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(runParallelCmd)
}
