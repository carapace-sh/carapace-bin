package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var profileStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start profiler on a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profileStartCmd).Standalone()

	profileStartCmd.Flags().String("clock-type", "", "specify the `TYPE` of clock used to report timestamps")
	profileStartCmd.Flags().String("sampling", "", "use sample profiling with `INTERVAL` microseconds between samples")
	profileStartCmd.Flags().Bool("streaming", false, "stream the profiling output to the specified file")
	profileStartCmd.Flags().String("user", "", "specify which `USER_ID` to profile")

	profileCmd.AddCommand(profileStartCmd)

	carapace.Gen(profileStartCmd).FlagCompletion(carapace.ActionMap{
		"clock-type": carapace.ActionValues("wall", "thread-cpu", "dual"),
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(profileStartCmd).PositionalCompletion(
		android.ActionPackages(),
		carapace.ActionFiles(),
	)
}
