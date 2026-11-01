package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var disableUntilUsedCmd = &cobra.Command{
	Use:   "disable-until-used",
	Short: "Disable the given package or component until used",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableUntilUsedCmd).Standalone()

	disableUntilUsedCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(disableUntilUsedCmd)

	carapace.Gen(disableUntilUsedCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(disableUntilUsedCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
