package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webhooksCmd = &cobra.Command{
	Use:     "webhooks",
	Aliases: []string{"webhook"},
	Short:   "Manage webhooks",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webhooksCmd).Standalone()

	rootCmd.AddCommand(webhooksCmd)
}
