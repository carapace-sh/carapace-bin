package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimental_runGraphCmd = &cobra.Command{
	Use:   "run-graph",
	Short: "Generate a graph of the execution order",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_runGraphCmd).Standalone()

	experimental_runGraphCmd.Flags().StringP("label", "l", "", "Label used in graph nodes (it could be either \"stack.name\" or \"stack.dir\"")
	experimental_runGraphCmd.Flags().StringP("outfile", "o", "", "Output .dot file")
	experimentalCmd.AddCommand(experimental_runGraphCmd)

	carapace.Gen(experimental_runGraphCmd).FlagCompletion(carapace.ActionMap{
		"outfile": carapace.ActionFiles(),
	})
}
