package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimental_evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "Eval expression",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_evalCmd).Standalone()

	experimental_evalCmd.Flags().Bool("as-json", false, "Outputs the result as a JSON value")
	experimental_evalCmd.Flags().StringSliceP("global", "g", nil, "set/override globals. eg.: --global name=<expr>")
	experimentalCmd.AddCommand(experimental_evalCmd)
}
