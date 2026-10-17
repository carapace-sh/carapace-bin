package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up the skill for installed agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillSetupCmd).Standalone()

	skillSetupCmd.Flags().String("agent", "", "limit to the agent NAME (repeatable)")
	skillSetupCmd.Flags().Bool("dry-run", false, "report what would change without changing anything")
	skillSetupCmd.Flags().Bool("force", false, "replace a skill named rex that Rex didn't install")
	skillSetupCmd.Flags().Bool("json", false, "print the result as JSON")
	skillSetupCmd.Flags().Bool("project", false, "use the current repository instead of your account")

	skillCmd.AddCommand(skillSetupCmd)

	carapace.Gen(skillSetupCmd).FlagCompletion(carapace.ActionMap{
		"agent": carapace.ActionValues("claude", "codex", "opencode", "pi"),
	})
}
