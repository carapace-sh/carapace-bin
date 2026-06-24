package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a repository with Scalar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_registerCmd).Standalone()

	scalar_registerCmd.Flags().Bool("maintenance", false, "enable automatic maintenance")
	scalarCmd.AddCommand(scalar_registerCmd)
}
