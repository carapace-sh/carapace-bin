package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugrebuildfncacheCmd = &cobra.Command{
	Use:    "debugrebuildfncache",
	Short:  "rebuild the fncache file",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugrebuildfncacheCmd).Standalone()

	debugrebuildfncacheCmd.Flags().Bool("only-data", false, "only look for wrong .d files (much faster)")
	rootCmd.AddCommand(debugrebuildfncacheCmd)
}
