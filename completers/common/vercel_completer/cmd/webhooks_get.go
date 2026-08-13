package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webhooks_getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"inspect"},
	Short:   "Displays information related to a webhook",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webhooks_getCmd).Standalone()

	webhooks_getCmd.Flags().String("format", "", "Output format")
	webhooks_getCmd.Flags().Bool("json", false, "Output as JSON")

	webhooksCmd.AddCommand(webhooks_getCmd)

	carapace.Gen(webhooks_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})

	carapace.Gen(webhooks_getCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
