package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:   "depends",
	Short: "visualize tox environment dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()

	dependsCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	dependsCmd.Flags().String("exit-and-dump-after", "", "")
	dependsCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	dependsCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	dependsCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	dependsCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	dependsCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	dependsCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	dependsCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	dependsCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	dependsCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	dependsCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	rootCmd.AddCommand(dependsCmd)
}
