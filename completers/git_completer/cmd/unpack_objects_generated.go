package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unpack_objectsCmd = &cobra.Command{
	Use:     "unpack-objects",
	Short:   "Unpack objects from a packed archive",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(unpack_objectsCmd).Standalone()

	rootCmd.AddCommand(unpack_objectsCmd)
}
