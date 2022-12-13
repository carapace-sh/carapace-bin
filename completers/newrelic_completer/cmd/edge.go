package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edgeCmd = &cobra.Command{
	Use:   "edge",
	Short: "Interact with New Relic Edge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edgeCmd).Standalone()
	rootCmd.AddCommand(edgeCmd)
}
