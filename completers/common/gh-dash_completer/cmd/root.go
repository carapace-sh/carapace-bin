package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-dash",
	Short: "A beautiful CLI dashboard for GitHub",
	Long:  "https://github.com/dlvhdr/gh-dash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "c", "", "use this configuration file (default is $XDG_CONFIG_HOME/gh-dash/config.yml)")
	rootCmd.Flags().Bool("debug", false, "passing this flag will allow writing debug output to debug.log")
	rootCmd.Flags().BoolP("help", "h", false, "help for gh-dash")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
