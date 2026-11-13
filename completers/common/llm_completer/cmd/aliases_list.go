package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliases_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliases_listCmd).Standalone()

	aliases_listCmd.Flags().Bool("json", false, "Output as JSON")
	aliasesCmd.AddCommand(aliases_listCmd)
}
