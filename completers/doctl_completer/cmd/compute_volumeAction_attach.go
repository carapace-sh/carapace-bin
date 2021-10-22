package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volumeAction_attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach a volume to a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volumeAction_attachCmd).Standalone()
	compute_volumeAction_attachCmd.Flags().Bool("wait", false, "Wait for volume to attach")
	compute_volumeActionCmd.AddCommand(compute_volumeAction_attachCmd)
}
