package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dirtyCmd = &cobra.Command{
	Use:   "dirty",
	Short: "Mark files as dirty so they show up in `lore status` and get picked up by `lore stage` (no content is read or staged)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dirtyCmd).Standalone()

	helpCmd.AddCommand(help_dirtyCmd)
}
