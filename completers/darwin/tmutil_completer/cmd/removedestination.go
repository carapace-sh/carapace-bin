package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removedestinationCmd)
}

var removedestinationCmd = &cobra.Command{
	Use:   "removedestination",
	Short: "remove a destination from configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removedestinationCmd).Standalone()
}