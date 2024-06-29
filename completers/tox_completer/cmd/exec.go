package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute an arbitrary command within a tox environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().StringS("e", "e", "", "environment to run (default: py)")
	execCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	execCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	execCmd.Flags().String("exit-and-dump-after", "", "")
	execCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 2416698050)")
	execCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	execCmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	execCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	execCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	execCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	execCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	execCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	execCmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")
	execCmd.Flags().StringP("pkg-only,", "b", "", "")
	execCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	execCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	execCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	execCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	execCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	execCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: )")
	execCmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")
	execCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	execCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	execCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(execCmd)
}
