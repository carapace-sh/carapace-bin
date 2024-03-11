package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var contributionsCmd = &cobra.Command{
	Use:     "contributions",
	Short:   "Summarise contributions to Homebrew repositories",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(contributionsCmd).Standalone()

	contributionsCmd.Flags().Bool("csv", false, "Print a CSV of contributions across repositories over the time period.")
	contributionsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	contributionsCmd.Flags().Bool("from", false, "Date (ISO-8601 format) to start searching contributions. Omitting this flag searches the last year.")
	contributionsCmd.Flags().Bool("help", false, "Show this message.")
	contributionsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	contributionsCmd.Flags().Bool("repositories", false, "Specify a comma-separated list of repositories to search. Supported repositories: `brew`, `core`, `cask`, `aliases`, `autoupdate`, `bundle`, `command-not-found`, `test-bot`, `services`, `cask-fonts` and `cask-versions`. Omitting this flag, or specifying `--repositories=primary`, searches only the main repositories: brew,core,cask. Specifying `--repositories=all`, searches all repositories. ")
	contributionsCmd.Flags().Bool("to", false, "Date (ISO-8601 format) to stop searching contributions.")
	contributionsCmd.Flags().Bool("user", false, "Specify a comma-separated list of GitHub usernames or email addresses to find contributions from. Omitting this flag searches maintainers.")
	contributionsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(contributionsCmd)
}
