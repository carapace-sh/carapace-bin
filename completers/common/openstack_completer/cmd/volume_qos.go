package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_qosCmd = &cobra.Command{
	Use:   "qos",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_qosCmd).Standalone()

	volumeCmd.AddCommand(volume_qosCmd)
}
