package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "show tox configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().StringS("e", "e", "", "enumerate (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	configCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	configCmd.Flags().Bool("core", false, "show core options (by default is hidden unless -e ALL is passed) (default: False)")
	configCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	configCmd.Flags().String("exit-and-dump-after", "", "")
	configCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	configCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	configCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	configCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	configCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	configCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	configCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	configCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	configCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	configCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	configCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	configCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: )")
	configCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	configCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	configCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(configCmd)
}
