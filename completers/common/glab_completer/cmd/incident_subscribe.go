package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var incident_subscribeCmd = &cobra.Command{
	Use:     "subscribe <id>",
	Short:   "Subscribe to an incident.",
	Aliases: []string{"sub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_subscribeCmd).Standalone()

	incidentCmd.AddCommand(incident_subscribeCmd)

	// TODO positional completion
}
