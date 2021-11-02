package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volumeAction_detachCmd = &cobra.Command{
	Use:   "detach",
	Short: "Detach a volume from a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volumeAction_detachCmd).Standalone()
	compute_volumeAction_detachCmd.Flags().Bool("wait", false, "Wait for volume to detach")
	compute_volumeActionCmd.AddCommand(compute_volumeAction_detachCmd)
}
