package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_tag_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a tag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_tag_deleteCmd).Standalone()
	compute_tag_deleteCmd.Flags().BoolP("force", "f", false, "Delete tag without confirmation prompt")
	compute_tagCmd.AddCommand(compute_tag_deleteCmd)
}
