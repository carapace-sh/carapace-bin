package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugdagCmd = &cobra.Command{
	Use:    "debugdag",
	Short:  "format the changelog or an index DAG as a concise textual description",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdagCmd).Standalone()

	debugdagCmd.Flags().BoolP("branches", "b", false, "annotate with branch names")
	debugdagCmd.Flags().Bool("dots", false, "use dots for runs")
	debugdagCmd.Flags().BoolP("spaces", "s", false, "separate elements by spaces")
	debugdagCmd.Flags().BoolP("tags", "t", false, "use tags as labels")
	rootCmd.AddCommand(debugdagCmd)
}
