package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_indexChangedPathsCmd = &cobra.Command{
	Use:   "index-changed-paths",
	Short: "Build changed-path index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_indexChangedPathsCmd).Standalone()

	debug_indexChangedPathsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_indexChangedPathsCmd.Flags().Uint32P("limit", "n", 4294967295, "Limit number of revisions to index")
	debugCmd.AddCommand(debug_indexChangedPathsCmd)
}