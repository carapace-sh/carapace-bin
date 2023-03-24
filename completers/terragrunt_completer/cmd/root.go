package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
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

	rootCmd.AddGroup(
		&cobra.Group{ID: "", Title: ""}, // dummy command to have different highlighting than terraform
		&cobra.Group{ID: "terragrunt", Title: "Terragrunt Commands"},
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("terraform"),
	)
}
