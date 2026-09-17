package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bin_pathsCmd = &cobra.Command{
	Use:   "bin-paths",
	Short: "List all the active runtime bin paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bin_pathsCmd).Standalone()

	rootCmd.AddCommand(bin_pathsCmd)
}
