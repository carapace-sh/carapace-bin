package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_untagCmd = &cobra.Command{
	Use:   "untag",
	Short: "Remove a tag from a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_untagCmd).Standalone()
	compute_droplet_untagCmd.Flags().StringSlice("tag-name", []string{}, "Tag name to remove from Droplet")
	compute_dropletCmd.AddCommand(compute_droplet_untagCmd)
}
