package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var list_statuslineCmd = &cobra.Command{
	Use:   "statusline",
	Short: "Single-line status for shell prompts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_statuslineCmd).Standalone()

	list_statuslineCmd.Flags().String("format", "", "Output format (table, json, claude-code)")
	list_statuslineCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	listCmd.AddCommand(list_statuslineCmd)

	carapace.Gen(list_statuslineCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json", "claude-code"),
	})
}
