package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var release_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List releases in a repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_listCmd).Standalone()
	releaseCmd.AddCommand(release_listCmd)
}
