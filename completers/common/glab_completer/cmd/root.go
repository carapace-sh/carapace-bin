package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glab <command> <subcommand> [flags]",
	Short: "A GitLab CLI tool.",
	Long:  "https://gitlab.com/gitlab-org/cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("help", false, "Show help for this command.")
	rootCmd.Flags().BoolP("version", "v", false, "show glab version information")

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if aliases, err := action.LoadAliases(); err == nil {
			for key, value := range aliases {
				cmd.Root().AddCommand(&cobra.Command{Use: key, Short: value, Run: func(cmd *cobra.Command, args []string) {}})
			}
		}
	})
}
