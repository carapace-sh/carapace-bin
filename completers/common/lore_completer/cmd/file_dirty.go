package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dirtyCmd = &cobra.Command{
	Use:   "dirty",
	Short: "Mark files as dirty so they show up in `lore status` and get picked up by directory-scoped `lore stage` (no content is read or staged)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dirtyCmd).Standalone()

	file_dirtyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_dirtyCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	fileCmd.AddCommand(file_dirtyCmd)
}
