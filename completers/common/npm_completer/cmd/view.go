package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:     "view",
	Short:   "View registry info",
	Aliases: []string{"info", "show", "v"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewCmd).Standalone()
	viewCmd.Flags().Bool("json", false, "output as json")
	addWorkspaceFlags(viewCmd)

	rootCmd.AddCommand(viewCmd)

	carapace.Gen(viewCmd).PositionalCompletion(
		action.ActionPackages(viewCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return npm.ActionPackumentFields(npm.PackageOpts{
				Registry: viewCmd.Flag("registry").Value.String(),
				Package:  c.Args[0],
			})
		}),
	)
}
