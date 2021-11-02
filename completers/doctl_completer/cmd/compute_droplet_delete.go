package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_deleteCmd).Standalone()
	compute_droplet_deleteCmd.Flags().BoolP("force", "f", false, "Delete the Droplet without a confirmation prompt")
	compute_droplet_deleteCmd.Flags().String("tag-name", "", "Tag name")
	compute_dropletCmd.AddCommand(compute_droplet_deleteCmd)
}
