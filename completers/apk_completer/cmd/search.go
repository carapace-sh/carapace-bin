package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search for packages by name or description",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "querying package information",
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("all", "a", false, "Print all matching package versions")
	searchCmd.Flags().BoolP("description", "d", false, "Also search for PATTERN in the package description")
	searchCmd.Flags().BoolP("exact", "e", false, "Match package names exactly")
	searchCmd.Flags().Bool("has-origin", false, "Match by package origin")
	searchCmd.Flags().BoolP("origin", "o", false, "Print base package name")
	searchCmd.Flags().BoolP("rdepends", "r", false, "Print reverse dependencies")
	searchCmd.Flags().BoolS("x", "x", false, "Match package names exactly")
	common.AddGlobalFlags(searchCmd)
	common.AddSourceFlags(searchCmd)
	rootCmd.AddCommand(searchCmd)
}
