package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var incident_unsubscribeCmd = &cobra.Command{
	Use:     "unsubscribe <id>",
	Short:   "Unsubscribe from an incident.",
	Aliases: []string{"unsub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_unsubscribeCmd).Standalone()

	incidentCmd.AddCommand(incident_unsubscribeCmd)

	// TODO positional completion
}
