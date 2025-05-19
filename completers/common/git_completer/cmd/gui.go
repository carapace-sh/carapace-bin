package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guiCmd = &cobra.Command{
	Use:     "gui",
	Short:   "A portable graphical interface to Git",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(guiCmd).Standalone()

	rootCmd.AddCommand(guiCmd)
}
