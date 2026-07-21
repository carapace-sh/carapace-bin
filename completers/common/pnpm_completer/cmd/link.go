package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Aliases: []string{"ln"},
	Short:   "Connect the local project to another one",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.Batch(
			pnpm.ActionDependencyNames(),
			carapace.ActionDirectories(),
		).ToA(),
	)
}
