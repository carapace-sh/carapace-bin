package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var profileStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop profiler on a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profileStopCmd).Standalone()

	profileStopCmd.Flags().String("user", "", "specify which `USER_ID` to profile")

	profileCmd.AddCommand(profileStopCmd)

	carapace.Gen(profileStopCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(profileStopCmd).PositionalCompletion(
		android.ActionPackages(),
	)
}
