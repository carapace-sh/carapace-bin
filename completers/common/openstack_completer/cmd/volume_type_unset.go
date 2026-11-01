package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var volume_type_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset volume type properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_type_unsetCmd).Standalone()

	volume_type_unsetCmd.Flags().Bool("encryption-type", false, "Remove the encryption type for this volume type (admin only)")
	volume_type_unsetCmd.Flags().String("project", "", "Removes volume type access to project (name or ID) (admin only)")
	volume_type_unsetCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	volume_type_unsetCmd.Flags().String("property", "", "Remove a property from this volume type (repeat option to remove multiple properties)")
	volume_typeCmd.AddCommand(volume_type_unsetCmd)

	carapace.Gen(volume_type_unsetCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(volume_type_unsetCmd),
		"project-domain": action.ActionDomains(volume_type_unsetCmd),
	})

	carapace.Gen(volume_type_unsetCmd).PositionalCompletion(action.ActionVolumeTypes(volume_type_unsetCmd))
}
