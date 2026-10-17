package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var skillRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the skill and hooks from agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillRemoveCmd).Standalone()

	common.AddSkillAgentFlags(skillRemoveCmd)

	skillCmd.AddCommand(skillRemoveCmd)

	carapace.Gen(skillRemoveCmd).FlagCompletion(carapace.ActionMap{
		"agent": carapace.ActionValues("claude", "codex", "opencode", "pi"),
	})
}
