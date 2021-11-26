package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for the given name in all packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_searchCmd).Standalone()
	pkg_searchCmd.Flags().BoolP("verbose", "v", false, "Show more information")
	pkgCmd.AddCommand(pkg_searchCmd)
}
