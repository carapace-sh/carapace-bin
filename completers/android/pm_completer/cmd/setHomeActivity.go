package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var setHomeActivityCmd = &cobra.Command{
	Use:   "set-home-activity",
	Short: "Set the default home activity (aka launcher)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setHomeActivityCmd).Standalone()

	setHomeActivityCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(setHomeActivityCmd)

	carapace.Gen(setHomeActivityCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(setHomeActivityCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
