package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Interact with New Relic entities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityCmd).Standalone()
	rootCmd.AddCommand(entityCmd)
}
