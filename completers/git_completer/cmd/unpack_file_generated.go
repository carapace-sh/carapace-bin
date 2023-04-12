package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unpack_fileCmd = &cobra.Command{
	Use:     "unpack-file",
	Short:   "Creates a temporary file with a blob's contents",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(unpack_fileCmd).Standalone()

	rootCmd.AddCommand(unpack_fileCmd)
}
