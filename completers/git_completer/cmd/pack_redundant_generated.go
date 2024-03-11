package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pack_redundantCmd = &cobra.Command{
	Use:     "pack-redundant",
	Short:   "Find redundant pack files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(pack_redundantCmd).Standalone()

	rootCmd.AddCommand(pack_redundantCmd)
}
