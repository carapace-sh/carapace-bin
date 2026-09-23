package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var grantCmd = &cobra.Command{
	Use:   "grant",
	Short: "Grant a permission to an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grantCmd).Standalone()

	grantCmd.Flags().String("user", "", "specify the `USER_ID` for which the operation needs to be performed")

	rootCmd.AddCommand(grantCmd)

	carapace.Gen(grantCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(grantCmd).PositionalCompletion(
		android.ActionPackages(),
		android.ActionPermissions(),
	)
}
