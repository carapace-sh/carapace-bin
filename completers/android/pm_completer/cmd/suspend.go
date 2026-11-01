package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "Suspends the specified packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendCmd).Standalone()

	suspendCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(suspendCmd)

	carapace.Gen(suspendCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(suspendCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
