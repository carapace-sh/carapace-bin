package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var cloud_version_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_updateCmd).Standalone()

	cloud_version_updateCmd.Flags().StringP("description", "d", "", "A description for this version")
	cloud_versionCmd.AddCommand(cloud_version_updateCmd)

	carapace.Gen(cloud_version_updateCmd).PositionalCompletion(
		vagrant.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
