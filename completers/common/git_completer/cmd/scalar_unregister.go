package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_unregisterCmd = &cobra.Command{
	Use:   "unregister",
	Short: "Unregister a repository from Scalar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_unregisterCmd).Standalone()

	scalarCmd.AddCommand(scalar_unregisterCmd)
}
