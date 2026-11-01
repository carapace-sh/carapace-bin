package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove the given package name from the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().BoolP("keep-data", "k", false, "keep the data and cache directories around after package removal")
	uninstallCmd.Flags().String("user", "", "remove the app from the given `USER_ID`")
	uninstallCmd.Flags().String("versionCode", "", "only uninstall if the app has the given version `code`")

	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
