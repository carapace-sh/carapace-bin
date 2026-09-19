package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugdirstateCmd = &cobra.Command{
	Use:    "debugdirstate",
	Short:  "show the contents of the current dirstate",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdirstateCmd).Standalone()

	debugdirstateCmd.Flags().Bool("all", false, "display dirstate-v2 tree nodes that would not exist in v1")
	debugdirstateCmd.Flags().Bool("dates", false, "display the saved mtime")
	debugdirstateCmd.Flags().Bool("datesort", false, "sort by saved mtime")
	debugdirstateCmd.Flags().Bool("docket", false, "display the docket (metadata file) instead")
	debugdirstateCmd.Flags().Bool("nodates", false, "do not display the saved mtime (DEPRECATED)")
	rootCmd.AddCommand(debugdirstateCmd)
}
