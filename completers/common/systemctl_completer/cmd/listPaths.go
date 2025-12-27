package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listPathsCmd = &cobra.Command{
	Use:     "list-paths",
	Short:   "List path units currently in memory, ordered by path",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPathsCmd).Standalone()

	rootCmd.AddCommand(listPathsCmd)
}
