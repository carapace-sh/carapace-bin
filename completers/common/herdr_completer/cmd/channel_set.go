package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var channel_setCmd = &cobra.Command{
	Use:   "set <stable|preview>",
	Short: "Choose the update channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(channel_setCmd).Standalone()

	channelCmd.AddCommand(channel_setCmd)

	carapace.Gen(channel_setCmd).PositionalCompletion(
		carapace.ActionValues("stable", "preview"),
	)
}
