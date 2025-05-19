package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var cloud_version_releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Releases a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_releaseCmd).Standalone()

	cloud_version_releaseCmd.Flags().BoolP("force", "f", false, "Release without confirmation")
	cloud_version_releaseCmd.Flags().Bool("no-force", false, "Do not release without confirmation")
	cloud_versionCmd.AddCommand(cloud_version_releaseCmd)

	carapace.Gen(cloud_version_releaseCmd).PositionalCompletion(
		vagrant.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
