package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Install Rex's skill and hooks into agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillSetupCmd).Standalone()

	skillSetupCmd.Flags().StringArray("agent", nil, "use only AGENT, which is claude, codex, opencode or pi (repeatable)")
	skillSetupCmd.Flags().Bool("dry-run", false, "report what would change without changing anything")
	skillSetupCmd.Flags().Bool("force", false, "replace a skill named rex that Rex did not install")
	skillSetupCmd.Flags().Bool("json", false, "print the result as JSON")
	skillSetupCmd.Flags().Bool("project", false, "set up the current repository instead of your account")
	skillSetupCmd.Flags().Bool("session-naming", false, "tell agents to name the session after their task")

	skillCmd.AddCommand(skillSetupCmd)

	carapace.Gen(skillSetupCmd).FlagCompletion(carapace.ActionMap{
		"agent": carapace.ActionValues("claude", "codex", "opencode", "pi"),
	})
}
