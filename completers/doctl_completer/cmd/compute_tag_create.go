package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_tag_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a tag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_tag_createCmd).Standalone()
	compute_tagCmd.AddCommand(compute_tag_createCmd)
}
