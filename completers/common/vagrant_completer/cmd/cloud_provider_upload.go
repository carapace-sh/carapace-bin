package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var cloud_provider_uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a box file to Vagrant Cloud for a specific provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_provider_uploadCmd).Standalone()

	cloud_provider_uploadCmd.Flags().BoolP("direct", "D", false, "Upload asset directly to backend storage")
	cloud_provider_uploadCmd.Flags().Bool("no-direct", false, "Do not upload asset directly to backend storage")
	cloud_providerCmd.AddCommand(cloud_provider_uploadCmd)

	carapace.Gen(cloud_provider_uploadCmd).PositionalCompletion(
		vagrant.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxProviders(c.Args[0])
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxVersions(c.Args[0], c.Args[1])
		}),
		carapace.ActionFiles(".box", ".json"),
	)
}
