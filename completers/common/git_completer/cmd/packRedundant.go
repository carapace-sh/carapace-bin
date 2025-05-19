package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packRedundantCmd = &cobra.Command{
	Use:     "pack-redundant",
	Short:   "Find redundant pack files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(packRedundantCmd).Standalone()

	packRedundantCmd.Flags().Bool("all", false, "processes all packs")
	packRedundantCmd.Flags().Bool("alt-odb", false, "don’t require objects from alternate odb to be present")
	packRedundantCmd.Flags().Bool("verbose", false, "outputs some statistics to stderr")
	rootCmd.AddCommand(packRedundantCmd)
}
