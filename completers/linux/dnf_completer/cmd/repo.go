package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo <subcommand> [options] [<repo-spec>...]",
	Short: "manage repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var repoListCmd = &cobra.Command{
	Use:   "list [repo-spec]...",
	Short: "list available repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var repoInfoCmd = &cobra.Command{
	Use:   "info [repo-spec]...",
	Short: "show detailed info about repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()

	repoCmd.AddCommand(repoListCmd)
	repoCmd.AddCommand(repoInfoCmd)

	for _, subcmd := range []*cobra.Command{repoListCmd, repoInfoCmd} {
		subcmd.Flags().Bool("all", false, "show information about all known repositories")
		subcmd.Flags().Bool("enabled", false, "show information about enabled repositories")
		subcmd.Flags().Bool("disabled", false, "show information about disabled repositories")
		subcmd.Flags().Bool("json", false, "show information in JSON format")
	}

	rootCmd.AddCommand(repoCmd)
}
