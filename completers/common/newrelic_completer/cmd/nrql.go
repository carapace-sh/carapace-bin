package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nrqlCmd = &cobra.Command{
	Use:   "nrql",
	Short: "Commands for interacting with the New Relic Database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nrqlCmd).Standalone()
	rootCmd.AddCommand(nrqlCmd)
}
