package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List captured threads, newest first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_listCmd).Standalone()

	ai_session_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ai_session_listCmd.Flags().String("style", "auto", "How output is rendered")
	ai_sessionCmd.AddCommand(ai_session_listCmd)

	carapace.Gen(ai_session_listCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
