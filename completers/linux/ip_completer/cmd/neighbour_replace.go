package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neighbour_replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "add new or change existing neighbour entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbour_replaceCmd).Standalone()

	neighbourCmd.AddCommand(neighbour_replaceCmd)
}
