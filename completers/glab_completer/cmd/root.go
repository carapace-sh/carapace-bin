package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glab",
	Short: "A GitLab CLI Tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.Flags().BoolP("version", "v", false, "show glab version information")

	carapace.Gen(rootCmd)
	addAliasCompletion(rootCmd)
}

func addAliasCompletion(cmd *cobra.Command) {
	if c, _, err := cmd.Find([]string{"_carapace"}); err == nil {
		c.Annotations = map[string]string{
			"skipAuthCheck": "true",
		}
		c.PreRun = func(cmd *cobra.Command, args []string) {
			if aliases, err := action.LoadAliases(); err == nil {
				for key, value := range aliases {
					cmd.Root().AddCommand(&cobra.Command{Use: key, Short: value, Run: func(cmd *cobra.Command, args []string) {}})
				}
			}
		}
	}
}
