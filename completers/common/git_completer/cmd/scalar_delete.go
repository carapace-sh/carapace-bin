package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing Scalar enlistment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_deleteCmd).Standalone()

	scalarCmd.AddCommand(scalar_deleteCmd)
}
