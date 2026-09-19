package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugStableTailSortLeapsCmd = &cobra.Command{
	Use:    "debug::stable-tail-sort-leaps",
	Short:  "REV",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugStableTailSortLeapsCmd).Standalone()

	debugStableTailSortLeapsCmd.Flags().BoolP("specific", "s", false, "restrict to specific leaps")
	debugStableTailSortLeapsCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugStableTailSortLeapsCmd)
}
