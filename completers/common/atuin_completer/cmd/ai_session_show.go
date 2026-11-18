package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a thread and its messages. Accepts an id, a harness id or `latest`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_showCmd).Standalone()

	ai_session_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ai_session_showCmd.Flags().String("style", "auto", "How output is rendered")
	ai_sessionCmd.AddCommand(ai_session_showCmd)

	carapace.Gen(ai_session_showCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
