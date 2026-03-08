package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoListCmd = &cobra.Command{
	Use:   "list [repo-spec]...",
	Short: "list available repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoListCmd).Standalone()

	repoListCmd.Flags().Bool("all", false, "show information about all known repositories")
	repoListCmd.Flags().Bool("disabled", false, "show information about disabled repositories")
	repoListCmd.Flags().Bool("enabled", false, "show information about enabled repositories")
	repoListCmd.Flags().Bool("json", false, "show information in JSON format")

	repoCmd.AddCommand(repoListCmd)
}
