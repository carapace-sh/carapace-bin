package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove Rex's skill and hooks from agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillRemoveCmd).Standalone()

	skillRemoveCmd.Flags().StringArray("agent", nil, "use only AGENT, which is claude, codex, opencode or pi (repeatable)")
	skillRemoveCmd.Flags().Bool("dry-run", false, "report what would change without changing anything")
	skillRemoveCmd.Flags().Bool("json", false, "print the result as JSON")
	skillRemoveCmd.Flags().Bool("project", false, "remove from the current repository instead of your account")

	skillCmd.AddCommand(skillRemoveCmd)

	carapace.Gen(skillRemoveCmd).FlagCompletion(carapace.ActionMap{
		"agent": carapace.ActionValues("claude", "codex", "opencode", "pi"),
	})
}
