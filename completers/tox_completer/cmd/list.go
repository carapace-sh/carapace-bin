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

	listCmd.Flags().BoolS("d", "d", false, "list just default envs (default: False)")
	listCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	listCmd.Flags().String("exit-and-dump-after", "", "")
	listCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	listCmd.Flags().Bool("no-desc", false, "do not show description (default: False)")
	listCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	listCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	listCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	listCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	listCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	listCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	listCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: )")
	listCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	listCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	listCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(listCmd)
}
