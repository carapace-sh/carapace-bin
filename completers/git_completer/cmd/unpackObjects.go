package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpackObjectsCmd = &cobra.Command{
	Use:     "unpack-objects",
	Short:   "Unpack objects from a packed archive",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(unpackObjectsCmd).Standalone()

	unpackObjectsCmd.Flags().String("max-input-size", "", "die, if the pack is larger than <size>")
	unpackObjectsCmd.Flags().BoolS("n", "n", false, "dry run")
	unpackObjectsCmd.Flags().BoolS("q", "q", false, "suppress percentage progress")
	unpackObjectsCmd.Flags().BoolS("r", "r", false, "try to recover objects  when unpacking a corrupt packfile")
	unpackObjectsCmd.Flags().Bool("strict", false, "don’t write objects with broken content or links")
	rootCmd.AddCommand(unpackObjectsCmd)
}
