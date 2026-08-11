package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dirtyCmd = &cobra.Command{
	Use:   "dirty",
	Short: "Mark files as dirty so they show up in `lore status` and get picked up by `lore stage` (no content is read or staged)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dirtyCmd).Standalone()

	dirtyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	dirtyCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	rootCmd.AddCommand(dirtyCmd)

	carapace.Gen(dirtyCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(dirtyCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
