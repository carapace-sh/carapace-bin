package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/am_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var startServiceCmd = &cobra.Command{
	Use:     "start-service",
	Aliases: []string{"startservice"},
	Short:   "Start a Service",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startServiceCmd).Standalone()

	startServiceCmd.Flags().String("user", "", "specify which `USER_ID` to run as")
	common.AddIntentFlags(startServiceCmd)

	rootCmd.AddCommand(startServiceCmd)

	carapace.Gen(startServiceCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(startServiceCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			android.ActionComponents(),
		).ToA(),
	)

}
