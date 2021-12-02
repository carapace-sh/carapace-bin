package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh",
	Short: "GitHub CLI",
	Long:  "https://cli.github.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.Flags().Bool("version", false, "Show gh version")

	carapace.Gen(rootCmd)
	addAliasAndExtensionCompletion(rootCmd)
}

func addAliasAndExtensionCompletion(cmd *cobra.Command) {
	if c, _, err := cmd.Find([]string{"_carapace"}); err == nil {
		c.Annotations = map[string]string{
			"skipAuthCheck": "true",
		}
		c.PreRun = func(cmd *cobra.Command, args []string) {
			if aliases, err := action.Aliases(); err == nil {
				for key, value := range aliases {
					cmd.Root().AddCommand(&cobra.Command{Use: key, Short: value, Run: func(cmd *cobra.Command, args []string) {}})
				}
			}
			if extensions, err := action.Extensions(); err == nil {
				for _, extension := range extensions {
					cmd.Root().AddCommand(&cobra.Command{Use: extension, Short: "extension", Run: func(cmd *cobra.Command, args []string) {}})
				}
			}
		}
	}
}
