package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var hideCmd = &cobra.Command{
	Use:   "hide",
	Short: "Hide the given package or component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hideCmd).Standalone()

	hideCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(hideCmd)

	carapace.Gen(hideCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(hideCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
