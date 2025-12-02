package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the branches in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_listCmd).Standalone()

	branch_listCmd.Flags().BoolP("all", "a", false, "Show all branches (not just active + 20 most recent)")
	branch_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_listCmd.Flags().BoolP("local", "l", false, "Show only local branches")
	branch_listCmd.Flags().Bool("no-ahead", false, "Don't calculate and show number of commits ahead of base (faster)")
	branch_listCmd.Flags().Bool("no-check", false, "Don't check if each branch merges cleanly into upstream")
	branch_listCmd.Flags().BoolP("remote", "r", false, "Show only remote branches")
	branch_listCmd.Flags().Bool("review", false, "Fetch and display review information (PRs, MRs, etc.)")
	branchCmd.AddCommand(branch_listCmd)
}
