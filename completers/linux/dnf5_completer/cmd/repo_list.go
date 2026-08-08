package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoListCmd = &cobra.Command{
	Use:   "list [options] [<repo-spec>...]",
	Short: "list repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoListCmd).Standalone()

	repoListCmd.Flags().Bool("all", false, "List all repositories")
	repoListCmd.Flags().Bool("disabled", false, "List disabled repositories")
	repoListCmd.Flags().Bool("enabled", false, "List enabled repositories (default)")
	repoListCmd.Flags().Bool("json", false, "Request json output format")

	repoCmd.AddCommand(repoListCmd)
}
