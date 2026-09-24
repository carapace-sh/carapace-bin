package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugStableTailSortCmd = &cobra.Command{
	Use:    "debug::stable-tail-sort",
	Short:  "REV",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugStableTailSortCmd).Standalone()

	debugStableTailSortCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugStableTailSortCmd)
}
