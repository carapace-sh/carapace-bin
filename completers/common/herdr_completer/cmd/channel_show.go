package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var channel_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Print the configured update channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(channel_showCmd).Standalone()

	channelCmd.AddCommand(channel_showCmd)
}
