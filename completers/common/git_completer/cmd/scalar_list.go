package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Scalar enlistments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_listCmd).Standalone()

	scalarCmd.AddCommand(scalar_listCmd)
}
