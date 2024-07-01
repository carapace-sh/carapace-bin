package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringS("e", "e", "", "enumerate (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	runCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	runCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	runCmd.Flags().String("exit-and-dump-after", "", "")
	runCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 525762999)")
	runCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	runCmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	runCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	runCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	runCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	runCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	runCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	runCmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")
	runCmd.Flags().StringP("pkg-only,", "b", "", "")
	runCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	runCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	runCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	runCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	runCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	runCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: )")
	runCmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")
	runCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	runCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	runCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(runCmd)
}
