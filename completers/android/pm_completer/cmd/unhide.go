package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var unhideCmd = &cobra.Command{
	Use:   "unhide",
	Short: "Unhide the given package or component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unhideCmd).Standalone()

	unhideCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(unhideCmd)

	carapace.Gen(unhideCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(unhideCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
