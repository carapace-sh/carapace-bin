package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mr_note_listCmd = &cobra.Command{
	Use:   "list [<id> | <branch>]",
	Short: "List merge request discussions. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_listCmd).Standalone()

	mr_note_listCmd.Flags().String("file", "", "Show only diff notes on this file path.")
	mr_note_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	mr_note_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	mr_note_listCmd.Flags().String("state", "", "Resolution state: all, resolved, unresolved.")
	mr_note_listCmd.Flags().StringP("type", "t", "", "Note type: all, general, diff, system.")
	mr_noteCmd.AddCommand(mr_note_listCmd)
}
