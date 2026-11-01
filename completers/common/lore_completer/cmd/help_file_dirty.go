package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_dirtyCmd = &cobra.Command{
	Use:   "dirty",
	Short: "Mark files as dirty so they show up in `lore status` and get picked up by directory-scoped `lore stage` (no content is read or staged)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_dirtyCmd).Standalone()

	help_fileCmd.AddCommand(help_file_dirtyCmd)
}
