package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdgraphCmd = &cobra.Command{
	Use:   "nerdgraph",
	Short: "Execute GraphQL requests to the NerdGraph API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdgraphCmd).Standalone()
	rootCmd.AddCommand(nerdgraphCmd)
}
