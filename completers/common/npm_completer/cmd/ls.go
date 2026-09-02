package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List installed packages",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()
	lsCmd.Flags().BoolP("all", "a", false, "show all outdated or installed packages")
	lsCmd.Flags().String("depth", "", "depth to ge when recursing packages")
	lsCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	lsCmd.Flags().String("include", "", "include dependency types")
	lsCmd.Flags().Bool("json", false, "output as json")
	lsCmd.Flags().Bool("link", false, "output only packages that are linked")
	lsCmd.Flags().BoolP("long", "l", false, "show extended information")
	lsCmd.Flags().StringArray("omit", nil, "omit dependency types")
	lsCmd.Flags().Bool("package-lock-only", false, "only use the package-lock.json")
	lsCmd.Flags().BoolP("parseable", "p", false, "output parseable results")
	lsCmd.Flags().Bool("unicode", false, "output unicode characters")
	addWorkspaceFlags(lsCmd)

	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"omit": carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(lsCmd).PositionalAnyCompletion(
		npm.ActionDependencyNames(),
	)
}
