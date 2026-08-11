package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_latest_listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_latest_listCmd).Standalone()

	branch_latest_listCmd.Flags().String("branch", "", "Branch to query")
	branch_latest_listCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_latestCmd.AddCommand(branch_latest_listCmd)
}
