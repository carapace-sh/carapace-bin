package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Stream server events as they occur",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()

	eventsCmd.Flags().Bool("all", false, "follow every server event, including clients joining, leaving, and changing")
	eventsCmd.Flags().Bool("json", false, "output events as newline-delimited JSON")

	eventsCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	rootCmd.AddCommand(eventsCmd)

	carapace.Gen(eventsCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(eventsCmd),
	})
}
