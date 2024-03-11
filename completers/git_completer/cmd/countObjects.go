package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var countObjectsCmd = &cobra.Command{
	Use:     "count-objects",
	Short:   "Count unpacked number of objects and their disk consumption",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(countObjectsCmd).Standalone()

	countObjectsCmd.Flags().BoolP("human-readable", "H", false, "print sizes in human readable format")
	countObjectsCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(countObjectsCmd)
}
