package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliases_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an alias for a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliases_setCmd).Standalone()

	aliases_setCmd.Flags().StringP("query", "q", "", "Set alias for model matching these strings")
	aliasesCmd.AddCommand(aliases_setCmd)
}
