package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var defaultStateCmd = &cobra.Command{
	Use:   "default-state",
	Short: "Reset the default state of the given package or component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultStateCmd).Standalone()

	defaultStateCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(defaultStateCmd)

	carapace.Gen(defaultStateCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(defaultStateCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
