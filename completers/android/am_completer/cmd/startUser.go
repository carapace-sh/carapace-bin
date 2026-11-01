package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var startUserCmd = &cobra.Command{
	Use:   "start-user",
	Short: "Start USER_ID in background if it is currently stopped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startUserCmd).Standalone()

	startUserCmd.Flags().String("display", "", "start the user visible in that `DISPLAY_ID`")
	startUserCmd.Flags().BoolP("wait", "w", false, "wait for start-user to complete and the user to be unlocked")

	rootCmd.AddCommand(startUserCmd)

	carapace.Gen(startUserCmd).PositionalCompletion(
		android.ActionUsers(),
	)
}
