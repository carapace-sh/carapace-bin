package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webhookCmd = &cobra.Command{
	Use:     "webhook",
	Short:   "Install the official webhook extension",
	GroupID: "extension",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webhookCmd).Standalone()

	rootCmd.AddCommand(webhookCmd)
}
