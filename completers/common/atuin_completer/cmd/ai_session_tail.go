package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Follow threads and messages as they are recorded (until interrupted)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_tailCmd).Standalone()

	ai_session_tailCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ai_session_tailCmd.Flags().String("style", "auto", "How output is rendered")
	ai_sessionCmd.AddCommand(ai_session_tailCmd)

	carapace.Gen(ai_session_tailCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
