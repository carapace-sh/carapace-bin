package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add a tag to a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_tagCmd).Standalone()
	compute_droplet_tagCmd.Flags().String("tag-name", "", "Tag name to use; can be a new or existing tag (required)")
	compute_dropletCmd.AddCommand(compute_droplet_tagCmd)
}
