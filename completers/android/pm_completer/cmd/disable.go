package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable the given package or component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableCmd).Standalone()

	disableCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(disableCmd)

	carapace.Gen(disableCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(disableCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
