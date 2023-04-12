package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "add, list, remove, update, and index chart repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()
	rootCmd.AddCommand(repoCmd)
}
