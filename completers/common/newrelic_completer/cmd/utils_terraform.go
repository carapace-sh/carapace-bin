package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utils_terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "Tools for working with Terraform",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utils_terraformCmd).Standalone()
	utilsCmd.AddCommand(utils_terraformCmd)
}
