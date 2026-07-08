package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ntable_changeCmd = &cobra.Command{
	Use:   "change",
	Short: "modify neighbour table parameters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ntable_changeCmd).Standalone()

	ntableCmd.AddCommand(ntable_changeCmd)
}
