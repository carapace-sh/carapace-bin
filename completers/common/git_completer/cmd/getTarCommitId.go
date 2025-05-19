package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getTarCommitIdCmd = &cobra.Command{
	Use:     "get-tar-commit-id",
	Short:   "Extract commit ID from an archive created using git-archive",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(getTarCommitIdCmd).Standalone()

	rootCmd.AddCommand(getTarCommitIdCmd)
}
