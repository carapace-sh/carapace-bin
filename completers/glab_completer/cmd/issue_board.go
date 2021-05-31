package cmd

import (
	"github.com/spf13/cobra"
)

var issue_boardCmd = &cobra.Command{
	Use:   "board",
	Short: "Work with GitLab Issue Boards in the given project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issueCmd.AddCommand(issue_boardCmd)
}
