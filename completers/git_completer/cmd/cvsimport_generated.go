package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cvsimportCmd = &cobra.Command{
	Use:     "cvsimport",
	Short:   "Salvage your data out of another SCM people love to hate",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(cvsimportCmd).Standalone()

	rootCmd.AddCommand(cvsimportCmd)
}
