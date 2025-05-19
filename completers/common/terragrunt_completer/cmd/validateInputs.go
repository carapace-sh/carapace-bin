package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var validateInputsCmd = &cobra.Command{
	Use:     "validate-inputs",
	Short:   "Checks if the terragrunt configured inputs align with the terraform defined variables",
	GroupID: "terragrunt",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateInputsCmd).Standalone()

	rootCmd.AddCommand(validateInputsCmd)
}
