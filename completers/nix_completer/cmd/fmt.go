package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:     "fmt",
	Short:   "reformat your code in the standard style",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	rootCmd.AddCommand(fmtCmd)

	addEvaluationFlags(fmtCmd)
	addFlakeFlags(fmtCmd)
	addLoggingFlags(fmtCmd)

	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
