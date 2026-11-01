package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to pane render updates (viewport and scrollback)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_subscribeCmd).Standalone()

	helpCmd.AddCommand(help_subscribeCmd)
}
