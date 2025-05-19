package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/fury"
	"github.com/spf13/cobra"
)

var beta_downloadCmd = &cobra.Command{
	Use:   "download PACKAGE@VERSION",
	Short: "Download a package to the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beta_downloadCmd).Standalone()

	betaCmd.AddCommand(beta_downloadCmd)

	carapace.Gen(betaCmd).PositionalCompletion(
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
