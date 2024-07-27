package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var channelListCmd = &cobra.Command{
	Use:   "channel:list",
	Short: "List all registered private broadcast channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(channelListCmd).Standalone()

	rootCmd.AddCommand(channelListCmd)
}
