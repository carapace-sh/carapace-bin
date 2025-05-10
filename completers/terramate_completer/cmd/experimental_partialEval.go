package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimental_partialEvalCmd = &cobra.Command{
	Use:   "partial-eval",
	Short: "Partial evaluate the expressions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_partialEvalCmd).Standalone()

	experimental_partialEvalCmd.Flags().StringSliceP("global", "g", nil, "set/override globals. eg.: --global name=<expr>")
	experimentalCmd.AddCommand(experimental_partialEvalCmd)
}
