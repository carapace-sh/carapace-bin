package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/fury"
	"github.com/spf13/cobra"
)

var yankCmd = &cobra.Command{
	Use:   "yank PACKAGE@VERSION",
	Short: "Remove a package version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(yankCmd).Standalone()

	yankCmd.Flags().StringP("version", "v", "", "Version")
	rootCmd.AddCommand(yankCmd)

	carapace.Gen(yankCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return fury.ActionPackages().Suffix("@").MultiParts(":")
			default:
				return fury.ActionPackageVersions(c.Parts[0])
			}
		}),
	)
}
