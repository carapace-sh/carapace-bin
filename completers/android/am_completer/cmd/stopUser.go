package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var stopUserCmd = &cobra.Command{
	Use:   "stop-user",
	Short: "Stop execution of USER_ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopUserCmd).Standalone()

	stopUserCmd.Flags().BoolP("force", "f", false, "force stop, even if user has an unstoppable parent")
	stopUserCmd.Flags().BoolP("wait", "w", false, "wait for stop-user to complete")

	rootCmd.AddCommand(stopUserCmd)

	carapace.Gen(stopUserCmd).PositionalCompletion(
		android.ActionUsers(),
	)
}
