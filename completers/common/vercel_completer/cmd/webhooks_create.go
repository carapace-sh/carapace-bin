package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var webhooks_createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a new webhook",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webhooks_createCmd).Standalone()

	webhooks_createCmd.Flags().String("event", "", "Event type for the webhook")
	webhooks_createCmd.Flags().String("project", "", "Project name or ID")

	webhooksCmd.AddCommand(webhooks_createCmd)

	carapace.Gen(webhooks_createCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(webhooks_createCmd),
	})

	carapace.Gen(webhooks_createCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
