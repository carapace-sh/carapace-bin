package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var cloud_provider_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a provider entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_provider_updateCmd).Standalone()

	cloud_provider_updateCmd.Flags().StringP("checksum", "c", "", "Checksum of the box for this provider. --checksum-type option is required.")
	cloud_provider_updateCmd.Flags().StringP("checksum-type", "C", "", "Type of checksum used (md5, sha1, sha256, sha384, sha512). --checksum option is required.")
	cloud_providerCmd.AddCommand(cloud_provider_updateCmd)

	carapace.Gen(cloud_provider_updateCmd).FlagCompletion(carapace.ActionMap{
		"checksum-type": carapace.ActionValues("md5", "sha1", "sha256", "sha384", "sha512"),
	})

	carapace.Gen(cloud_provider_updateCmd).PositionalCompletion(
		vagrant.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxProviders(c.Args[0])
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionCloudBoxVersions(c.Args[0], c.Args[1])
		}),
	)
}
