package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var cloud_version_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_createCmd).Standalone()

	cloud_version_createCmd.Flags().StringP("description", "d", "", "A description for this version")
	cloud_versionCmd.AddCommand(cloud_version_createCmd)

	carapace.Gen(cloud_version_createCmd).PositionalCompletion(
		vagrant.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
