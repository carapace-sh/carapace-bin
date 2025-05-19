package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var cloud_provider_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a provider entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_provider_deleteCmd).Standalone()

	cloud_provider_deleteCmd.Flags().BoolP("force", "f", false, "Force deletion of box version provider without confirmation")
	cloud_provider_deleteCmd.Flags().Bool("no-force", false, "Do not force deletion of box version provider without confirmation")
	cloud_providerCmd.AddCommand(cloud_provider_deleteCmd)

	carapace.Gen(cloud_provider_deleteCmd).PositionalCompletion(
		vagrant.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxProviders(c.Args[0])
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxVersions(c.Args[0], c.Args[1])
		}),
	)
}
