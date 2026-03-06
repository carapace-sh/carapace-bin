package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List worktrees and their status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().Bool("branches", false, "Include branches without worktrees")
	listCmd.Flags().String("format", "", "Output format (table, json)")
	listCmd.Flags().Bool("full", false, "Show CI, diff analysis, and LLM summaries")
	listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	listCmd.Flags().Bool("progressive", false, "Show fast info immediately, update with slow info")
	listCmd.Flags().Bool("remotes", false, "Include remote branches")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
	})
}
