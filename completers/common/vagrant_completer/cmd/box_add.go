package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var box_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a box",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(box_addCmd).Standalone()

	box_addCmd.Flags().String("box-version", "", "Constrain version of the added box")
	box_addCmd.Flags().String("cacert", "", "CA certificate for SSL download")
	box_addCmd.Flags().String("capath", "", "CA certificate directory for SSL download")
	box_addCmd.Flags().String("cert", "", "A client SSL cert, if needed")
	box_addCmd.Flags().String("checksum", "", "Checksum for the box")
	box_addCmd.Flags().String("checksum-type", "", "Checksum type (md5, sha1, sha256)")
	box_addCmd.Flags().BoolP("clean", "c", false, "Clean any temporary download files")
	box_addCmd.Flags().BoolP("force", "f", false, "Overwrite an existing box if it exists")
	box_addCmd.Flags().Bool("insecure", false, "Do not validate SSL certificates")
	box_addCmd.Flags().Bool("location-trusted", false, "Trust 'Location' header from HTTP redirects and use the same credentials for subsequent urls as for the initial one")
	box_addCmd.Flags().String("name", "", "Name of the box")
	box_addCmd.Flags().String("provider", "", "Provider the box should satisfy")
	boxCmd.AddCommand(box_addCmd)

	carapace.Gen(box_addCmd).FlagCompletion(carapace.ActionMap{
		"cacert":        carapace.ActionFiles(),
		"capath":        carapace.ActionDirectories(),
		"cert":          carapace.ActionFiles(),
		"checksum-type": carapace.ActionValues("md5", "sha1", "sha256"),
		"provider":      vagrant.ActionProviders(),
	})

	carapace.Gen(box_addCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			provider := box_addCmd.Flag("provider").Value.String()
			return carapace.Batch(
				carapace.ActionFiles(".box", ".json"),
				vagrant.ActionCloudBoxSearch(provider).UnlessF(condition.CompletingPath),
			).ToA()
		}),
	)
}
