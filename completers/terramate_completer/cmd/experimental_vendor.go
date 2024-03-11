package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimental_vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Manages vendored Terraform modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_vendorCmd).Standalone()

	experimentalCmd.AddCommand(experimental_vendorCmd)
}
