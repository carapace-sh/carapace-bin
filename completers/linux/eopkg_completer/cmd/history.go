package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Short:   "History of eopkg operations",
	Aliases: []string{"hs"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.Flags().StringP("last", "l", "", "Output only the last n operations")
	historyCmd.Flags().BoolP("snapshot", "s", false, "Take snapshot of the current system")
	historyCmd.Flags().StringP("takeback", "t", "", "Takeback to a given operation number")

	rootCmd.AddCommand(historyCmd)
}
