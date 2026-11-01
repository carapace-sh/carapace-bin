package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke a permission from an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revokeCmd).Standalone()

	revokeCmd.Flags().Bool("all-permissions", false, "grant all missing runtime permissions")
	revokeCmd.Flags().String("user", "", "specify the `USER_ID` for which the operation needs to be performed")

	rootCmd.AddCommand(revokeCmd)

	carapace.Gen(revokeCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(revokeCmd).PositionalCompletion(
		android.ActionPackages(),
		android.ActionPermissions(),
	)
}
