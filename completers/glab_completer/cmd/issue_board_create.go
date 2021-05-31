package cmd

import (
	"github.com/spf13/cobra"
)

var issue_board_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project issue board.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issue_board_createCmd.Flags().StringP("name", "n", "", "The name of the new board")
	issue_boardCmd.AddCommand(issue_board_createCmd)
}
