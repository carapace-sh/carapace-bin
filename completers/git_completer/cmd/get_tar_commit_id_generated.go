package cmd

import (
	"github.com/spf13/cobra"
)

var get_tar_commit_idCmd = &cobra.Command{
	Use:     "get-tar-commit-id",
	Short:   "Extract commit ID from an archive created using git-archive",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {

	rootCmd.AddCommand(get_tar_commit_idCmd)
}
