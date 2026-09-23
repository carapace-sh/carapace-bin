package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var disableUserCmd = &cobra.Command{
	Use:   "disable-user",
	Short: "Disable the given package or component for the user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableUserCmd).Standalone()

	disableUserCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(disableUserCmd)

	carapace.Gen(disableUserCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(disableUserCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
