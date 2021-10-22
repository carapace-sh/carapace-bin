package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_volumeAction_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize the disk of a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volumeAction_resizeCmd).Standalone()
	compute_volumeAction_resizeCmd.Flags().String("region", "", "Volume region (required)")
	compute_volumeAction_resizeCmd.Flags().Int("size", 0, "New size in GiB (required)")
	compute_volumeAction_resizeCmd.Flags().Bool("wait", false, "Wait for volume to resize")
	compute_volumeActionCmd.AddCommand(compute_volumeAction_resizeCmd)

	carapace.Gen(compute_volumeAction_resizeCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
}
