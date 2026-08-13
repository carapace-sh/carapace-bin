package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webhooks_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Show all webhooks",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webhooks_lsCmd).Standalone()

	webhooks_lsCmd.Flags().String("format", "", "Output format")
	webhooks_lsCmd.Flags().Bool("json", false, "Output as JSON")

	webhooksCmd.AddCommand(webhooks_lsCmd)
}
