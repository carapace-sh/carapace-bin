package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volumeAction_detachByDropletIdCmd = &cobra.Command{
	Use:   "detach-by-droplet-id",
	Short: "(Deprecated) Detach a volume. Use `detach` instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volumeAction_detachByDropletIdCmd).Standalone()
	compute_volumeActionCmd.AddCommand(compute_volumeAction_detachByDropletIdCmd)
}
