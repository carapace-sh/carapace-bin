package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugindexstatsCmd = &cobra.Command{
	Use:    "debugindexstats",
	Short:  "show stats related to the changelog index",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugindexstatsCmd).Standalone()

	rootCmd.AddCommand(debugindexstatsCmd)
}
