package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// add_common_flags defines the flags on a tox subcommand that are common to all
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
