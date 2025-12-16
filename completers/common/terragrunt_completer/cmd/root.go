package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "terragrunt",
	Short:              "Terragrunt is a thin wrapper for Terraform",
	Long:               "https://terragrunt.gruntwork.io/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "", Title: ""}, // dummy command to have different highlighting than terraform
		&cobra.Group{ID: "terragrunt", Title: "Terragrunt Commands"},
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO support `TERRAGRUNT_TFPATH`
			if _, err := execlog.LookPath("terraform"); err == nil {
				return bridge.ActionCarapaceBin("terraform")
			}
			return bridge.ActionCarapaceBin("tofu")
		}),
	)
}
