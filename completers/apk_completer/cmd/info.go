package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Give detailed information about packages or repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "querying package information",
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolP("all", "a", false, "List all information known about the package")
	infoCmd.Flags().BoolP("contents", "L", false, "List files included in the package")
	infoCmd.Flags().BoolP("depends", "R", false, "List the dependencies of the package")
	infoCmd.Flags().BoolP("description", "d", false, "Print the package description")
	infoCmd.Flags().Bool("install-if", false, "List the package's install_if rule")
	infoCmd.Flags().BoolP("installed", "e", false, "Check package installed status")
	infoCmd.Flags().Bool("license", false, "Print the package SPDX license identifier")
	infoCmd.Flags().BoolP("provides", "P", false, "List what the package provides")
	infoCmd.Flags().BoolP("rdepends", "r", false, "List reverse dependencies of the package")
	infoCmd.Flags().Bool("replaces", false, "List the other packages for which this package is marked as a replacement")
	infoCmd.Flags().Bool("rinstall-if", false, "List other packages whose install_if rules refer to this package")
	infoCmd.Flags().BoolP("size", "s", false, "Print the package's installed size")
	infoCmd.Flags().BoolP("triggers", "t", false, "Print active triggers for the package")
	infoCmd.Flags().BoolP("webpage", "w", false, "Print the URL for the package's upstream webpage")
	infoCmd.Flags().BoolP("who-owns", "W", false, "Print the package which owns the specified file")
	common.AddGlobalFlags(infoCmd)
	common.AddSourceFlags(infoCmd)
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(infoCmd).FilterArgs(),
	)
}
