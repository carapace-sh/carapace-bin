package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hash_objectCmd = &cobra.Command{
	Use:     "hash-object",
	Short:   "Compute object ID and optionally creates a blob from a file",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(hash_objectCmd).Standalone()
	hash_objectCmd.Flags().Bool("literally", false, "just hash any random garbage to create corrupt objects for debugging Git")
	hash_objectCmd.Flags().Bool("no-filters", false, "store file as is without filters")
	hash_objectCmd.Flags().String("path", "", "process file as it were from this path")
	hash_objectCmd.Flags().Bool("stdin", false, "read the object from stdin")
	hash_objectCmd.Flags().Bool("stdin-paths", false, "read file names from stdin")
	hash_objectCmd.Flags().StringS("t", "t", "", "object type")
	hash_objectCmd.Flags().BoolS("w", "w", false, "write the object into the object database")
	rootCmd.AddCommand(hash_objectCmd)
}
