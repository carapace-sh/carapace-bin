package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:     "stats",
	Short:   "Show statistics about repositories and installations",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "miscellaneous",
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	rootCmd.AddCommand(statsCmd)
}
