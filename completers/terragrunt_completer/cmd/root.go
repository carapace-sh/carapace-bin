package cmd

import (
	"github.com/rsteube/carapace"
	terraform "github.com/rsteube/carapace-bin/completers/terraform_completer/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "terragrunt",
	Short:              "Terragrunt is a thin wrapper for Terraform",
	Long:               "https://terragrunt.gruntwork.io/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		terraform.ActionExecute(),
	)
}
