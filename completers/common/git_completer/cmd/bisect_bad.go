package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bisect_badCmd = &cobra.Command{
	Use:     "bad",
	Aliases: []string{"new"},
	Short:   "mark a known-bad revision",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_badCmd).Standalone()
	bisectCmd.AddCommand(bisect_badCmd)

	carapace.Gen(bisect_badCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(bisect_badCmd).DashAnyCompletion(
		carapace.ActionPositional(bisect_badCmd),
	)
}
