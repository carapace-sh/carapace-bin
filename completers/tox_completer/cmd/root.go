package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tox",
	Short: "automation project",
	Long:  "https://tox.wiki/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	rootCmd.Flags().String("colored", "", "should output be enriched with colors, default is yes unless TERM=dumb or NO_COLOR is defined. (default: no)")
	rootCmd.Flags().String("exit-and-dump-after", "", "")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().String("no-provision", "", "do not perform provision, but fail and if a path was provided write provision metadata as JSON to it (default: False)")
	rootCmd.Flags().Bool("no-recreate-provision", false, "if recreate is set do not recreate provision tox environment (default: False)")
	rootCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity (default: 0)")
	rootCmd.Flags().BoolP("recreate", "r", false, "recreate the tox environments (default: False)")
	rootCmd.Flags().String("root", "", "project root directory (if not specified will be the folder of the config file) (default: None)")
	rootCmd.Flags().String("runner", "", "the tox run engine to use when not explicitly stated in tox env configuration (default: virtualenv)")
	rootCmd.Flags().BoolP("verbose", "v", false, "increase verbosity (default: 2)")
	rootCmd.Flags().Bool("version", false, "show program's and plugins version number and exit")
	rootCmd.Flags().String("workdir", "", "tox working directory (if not specified will be the folder of the config file) (default: None)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"e":  tox.ActionEnvironments().UniqueList(","),
	})
}
