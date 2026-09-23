package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var unsuspendCmd = &cobra.Command{
	Use:   "unsuspend",
	Short: "Unsuspends the specified packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsuspendCmd).Standalone()

	unsuspendCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(unsuspendCmd)

	carapace.Gen(unsuspendCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(unsuspendCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
