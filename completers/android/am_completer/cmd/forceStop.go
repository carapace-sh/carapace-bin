package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var forceStopCmd = &cobra.Command{
	Use:   "force-stop",
	Short: "Completely stop the given application package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forceStopCmd).Standalone()

	forceStopCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(forceStopCmd)

	carapace.Gen(forceStopCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("all", "current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(forceStopCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
