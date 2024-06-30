package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	// common
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	rootCmd.Flags().Uint("exit-and-dump-after", 0, "dump tox threads after n seconds and exit the app - useful to debug when tox hangs, 0 means disabled (default: 0)")
	rootCmd.Flags().StringP("conf", "c", "", "configuration file/folder for tox (if not specified will discover one) (default: None)")
	rootCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")
	rootCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	rootCmd.Flags().String("runner", "virtualenv", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	rootCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	rootCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	rootCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	rootCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	rootCmd.Flags().StringArrayP("override", "x", []string{}, "configuration override(s), e.g., -x testenv:pypy3.ignore_errors=True (default: [])")
	rootCmd.Flags().BoolSliceP("quiet", "q", []bool{}, "decrease verbosity (default: 0)")
	rootCmd.Flags().BoolSliceP("verbose", "v", []bool{}, "increase verbosity (default: 2)")

	// specific to this subcommend
	rootCmd.Flags().StringArrayS("m", []string{}, "labels to evaluate (default: [])")

	rootCmd.AddCommand(listCmd)
}
