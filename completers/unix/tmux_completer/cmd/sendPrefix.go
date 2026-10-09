package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var sendPrefixCmd = &cobra.Command{
	Use:   "send-prefix",
	Short: "send the prefix key to a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendPrefixCmd).Standalone()

	sendPrefixCmd.Flags().BoolS("2", "2", false, "send secondary prefix key")
	sendPrefixCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(sendPrefixCmd)

	carapace.Gen(sendPrefixCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
