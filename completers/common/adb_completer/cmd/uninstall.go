package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "remove this app package from the device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().String("user", "", "remove the app for the given user")
	uninstallCmd.Flags().String("versionCode", "", "only remove the app if it has this version code")

	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"user": actionInstallUsers(),
	})

	carapace.Gen(uninstallCmd).PositionalCompletion(
		adb.ActionPackages(),
	)
}
