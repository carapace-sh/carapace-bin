package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/am_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var startForegroundServiceCmd = &cobra.Command{
	Use:     "start-foreground-service",
	Aliases: []string{"startforegroundservice", "start-fg-service"},
	Short:   "Start a foreground Service",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startForegroundServiceCmd).Standalone()

	startForegroundServiceCmd.Flags().String("user", "", "specify which `USER_ID` to run as")
	common.AddIntentFlags(startForegroundServiceCmd)

	rootCmd.AddCommand(startForegroundServiceCmd)

	carapace.Gen(startForegroundServiceCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(startForegroundServiceCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			android.ActionComponents(),
		).ToA(),
	)

}
