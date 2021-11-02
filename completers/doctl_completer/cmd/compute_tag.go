package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Display commands to manage tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_tagCmd).Standalone()
	computeCmd.AddCommand(compute_tagCmd)
}
