package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devenvCmd = &cobra.Command{
	Use:   "devenv",
	Short: "sets up a development environment at ENVDIR based on the tox configuration specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devenvCmd).Standalone()

	devenvCmd.Flags().StringS("e", "e", "", "environment to run (default: py)")
	devenvCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	devenvCmd.Flags().String("exit-and-dump-after", "", "")
	devenvCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 3925389505)")
	devenvCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	devenvCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	devenvCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	devenvCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	devenvCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	devenvCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	devenvCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	devenvCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	devenvCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	devenvCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	devenvCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	devenvCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: )")
	devenvCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	devenvCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	devenvCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(devenvCmd)
}
