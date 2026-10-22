package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var volume_qos_associateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate a QoS specification to a volume type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_qos_associateCmd).Standalone()

	volume_qosCmd.AddCommand(volume_qos_associateCmd)

	carapace.Gen(volume_qos_associateCmd).PositionalCompletion(
		action.ActionVolumeQosSpecs(),
		action.ActionVolumeTypes(),
	)
}
