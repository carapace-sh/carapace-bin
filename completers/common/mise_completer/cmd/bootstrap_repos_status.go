package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_repos_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the state of Git repositories from `[bootstrap.repos]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_repos_statusCmd).Standalone()

	bootstrap_repos_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_repos_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_repos_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if any repository differs from the config (all are still listed)")
	bootstrap_reposCmd.AddCommand(bootstrap_repos_statusCmd)
}
