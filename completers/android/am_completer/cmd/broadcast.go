package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/am_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "Send a broadcast Intent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(broadcastCmd).Standalone()

	broadcastCmd.Flags().Bool("allow-background-activity-starts", false, "the receiver may start activities even if in the background")
	broadcastCmd.Flags().Bool("async", false, "send without waiting for the completion of the receiver")
	broadcastCmd.Flags().String("receiver-permission", "", "require `PERMISSION` of the receiver")
	broadcastCmd.Flags().String("user", "", "specify which `USER_ID` to send to")
	common.AddIntentFlags(broadcastCmd)

	rootCmd.AddCommand(broadcastCmd)

	carapace.Gen(broadcastCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("all", "current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(broadcastCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			android.ActionComponents(),
		).ToA(),
	)

}
