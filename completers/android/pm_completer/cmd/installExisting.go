package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var installExistingCmd = &cobra.Command{
	Use:   "install-existing",
	Short: "Installs an existing application for a new user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installExistingCmd).Standalone()

	installExistingCmd.Flags().Bool("full", false, "install as a full app")
	installExistingCmd.Flags().Bool("instant", false, "install as an instant app")
	installExistingCmd.Flags().Bool("restrict-permissions", false, "don't whitelist restricted permissions")
	installExistingCmd.Flags().String("user", "", "install for the given `USER_ID`")
	installExistingCmd.Flags().Bool("wait", false, "wait until the package is installed")

	rootCmd.AddCommand(installExistingCmd)

	carapace.Gen(installExistingCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("all", "current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(installExistingCmd).PositionalCompletion(
		android.ActionPackages(),
	)
}
