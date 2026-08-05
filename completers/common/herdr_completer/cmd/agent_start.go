package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_startCmd = &cobra.Command{
	Use:   "start <name>",
	Short: "Start a supported interactive agent in an existing pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_startCmd).Standalone()

	agent_startCmd.Flags().String("kind", "", "Supported agent kind and canonical executable")
	agent_startCmd.Flags().String("pane", "", "Existing pane at an interactive shell prompt")
	agent_startCmd.Flags().String("timeout", "", "Wait for interactive readiness (default: 30000; max: 300000)")
	agent_startCmd.MarkFlagRequired("kind")
	agent_startCmd.MarkFlagRequired("pane")
	agentCmd.AddCommand(agent_startCmd)

	carapace.Gen(agent_startCmd).FlagCompletion(carapace.ActionMap{
		"kind": carapace.ActionValues("pi", "claude", "codex", "gemini", "cursor", "devin", "agy", "cline", "omp", "mastracode", "opencode", "copilot", "kimi", "kiro", "droid", "amp", "grok", "hermes", "kilo", "qodercli", "maki"),
		"pane": herdr.ActionPanes(herdr.PaneOpts{}),
	})
}
