package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_listCmd).Standalone()
	pkg_listCmd.Flags().StringP("output", "o", "list", "Defines the output format (valid: 'list', 'json')")
	pkg_listCmd.Flags().BoolP("verbose", "v", false, "Show more information")
	pkgCmd.AddCommand(pkg_listCmd)

	carapace.Gen(pkg_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("list", "json"),
	})
}
