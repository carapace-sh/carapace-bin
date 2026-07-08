package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neighbour_changeCmd = &cobra.Command{
	Use:   "change",
	Short: "change a neighbour entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbour_changeCmd).Standalone()

	neighbourCmd.AddCommand(neighbour_changeCmd)
}
