package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/am_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var stopServiceCmd = &cobra.Command{
	Use:     "stop-service",
	Aliases: []string{"stopservice"},
	Short:   "Stop a Service",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopServiceCmd).Standalone()

	stopServiceCmd.Flags().String("user", "", "specify which `USER_ID` to run as")
	common.AddIntentFlags(stopServiceCmd)

	rootCmd.AddCommand(stopServiceCmd)

	carapace.Gen(stopServiceCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(stopServiceCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			android.ActionComponents(),
		).ToA(),
	)

}
