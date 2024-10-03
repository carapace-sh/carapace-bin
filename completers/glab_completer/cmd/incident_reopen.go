package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var incident_reopenCmd = &cobra.Command{
	Use:     "reopen [<id> | <url>] [flags]",
	Short:   "Reopen a resolved incident.",
	Aliases: []string{"open"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_reopenCmd).Standalone()

	incidentCmd.AddCommand(incident_reopenCmd)

	// TODO positional completion
}
