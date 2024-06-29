package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quickstartCmd = &cobra.Command{
	Use:   "quickstart",
	Short: "Command line script to quickly create a tox config file for a Python project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quickstartCmd).Standalone()

	quickstartCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	quickstartCmd.Flags().String("exit-and-dump-after", "", "")
	quickstartCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	quickstartCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	quickstartCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	quickstartCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	quickstartCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	quickstartCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	quickstartCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	quickstartCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	quickstartCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	quickstartCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(quickstartCmd)
}
