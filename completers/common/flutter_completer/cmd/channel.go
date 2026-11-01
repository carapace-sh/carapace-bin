package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "List or switch Flutter channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(channelCmd).Standalone()

	channelCmd.Flags().BoolP("all", "a", false, "Include all the available branches (including local branches) when listing channels.")
	channelCmd.Flags().Bool("cache-artifacts", false, "After switching channels, download all required binary artifacts. This is the equivalent of running \"flutter precache\" with the \"--all-platforms\" flag.")
	channelCmd.Flags().BoolP("force", "f", false, "Force switch channels, potentially discarding local changes.")
	channelCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	channelCmd.Flags().Bool("no-all", false, "Include all the available branches (including local branches) when listing channels.")
	channelCmd.Flags().Bool("no-cache-artifacts", false, "After switching channels, download all required binary artifacts. This is the equivalent of running \"flutter precache\" with the \"--all-platforms\" flag.")
	rootCmd.AddCommand(channelCmd)

	carapace.Gen(channelCmd).PositionalCompletion(
		action.ActionChannels(),
	)
}
