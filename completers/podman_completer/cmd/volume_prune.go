package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_pruneCmd = &cobra.Command{
	Use:   "prune [options]",
	Short: "Remove all unused volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_pruneCmd).Standalone()

	volume_pruneCmd.Flags().StringSlice("filter", []string{}, "Provide filter values (e.g. 'label=<key>=<value>')")
	volume_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	volumeCmd.AddCommand(volume_pruneCmd)
}
