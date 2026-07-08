package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neighbour_getCmd = &cobra.Command{
	Use:   "get",
	Short: "lookup a neighbour entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbour_getCmd).Standalone()

	neighbourCmd.AddCommand(neighbour_getCmd)
}
