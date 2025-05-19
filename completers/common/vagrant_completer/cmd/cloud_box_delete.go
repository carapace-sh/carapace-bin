package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var cloud_box_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes box entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_box_deleteCmd).Standalone()

	cloud_box_deleteCmd.Flags().BoolP("force", "f", false, "Do not prompt for deletion confirmation")
	cloud_box_deleteCmd.Flags().Bool("no-force", false, "Prompt for deletion confirmation")
	cloud_boxCmd.AddCommand(cloud_box_deleteCmd)

	carapace.Gen(cloud_box_deleteCmd).PositionalCompletion(
		vagrant.ActionCloudBoxSearch(""),
	)
}
