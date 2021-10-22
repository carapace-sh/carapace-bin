package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_repository_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories for a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_repository_listCmd).Standalone()
	registry_repository_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `LatestTag`, `TagCount`, `UpdatedAt`")
	registry_repository_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	registry_repositoryCmd.AddCommand(registry_repository_listCmd)
}
