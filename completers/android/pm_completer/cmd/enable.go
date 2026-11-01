package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable the given package or component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableCmd).Standalone()

	enableCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(enableCmd)

	carapace.Gen(enableCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(enableCmd).PositionalCompletion(
		android.ActionComponents(),
	)
}
