package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/am_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start-activity",
	Aliases: []string{"start"},
	Short:   "Start an Activity",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	startCmd.Flags().String("activity-type", "", "the `ACTIVITY_TYPE` to launch the activity as")
	startCmd.Flags().BoolP("debug", "D", false, "enable debugging")
	startCmd.Flags().String("display", "", "the `DISPLAY_ID` to launch the activity into")
	startCmd.Flags().BoolP("force-stop", "S", false, "force stop the target app before starting the activity")
	startCmd.Flags().BoolP("native-debugging", "N", false, "enable native debugging")
	startCmd.Flags().IntP("repeat", "R", 0, "repeat the activity launch COUNT times")
	startCmd.Flags().StringP("start-profiler", "P", "", "start profiler and send results to `FILE`")
	startCmd.Flags().Bool("suspend", false, "debugged app suspend threads at startup (only with -D)")
	startCmd.Flags().Bool("track-allocation", false, "enable tracking of object allocations")
	startCmd.Flags().String("user", "", "specify which `USER_ID` to run as")
	startCmd.Flags().BoolP("wait", "W", false, "wait for launch to complete (initial display)")
	startCmd.Flags().String("windowing-mode", "", "the `WINDOWING_MODE` to launch the activity into")
	common.AddIntentFlags(startCmd)

	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).FlagCompletion(carapace.ActionMap{
		"activity-type":  carapace.ActionValues("standard", "home", "secondaryHome", "multiInstance"),
		"start-profiler": carapace.ActionFiles(),
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
		"windowing-mode": carapace.ActionValues("fullscreen", "multi-window", "picture-in-picture", "freeform", "split-screen"),
	})
	carapace.Gen(startCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			android.ActionComponents(),
		).ToA(),
	)

}
