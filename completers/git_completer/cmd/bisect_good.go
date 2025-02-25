package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bisect_goodCmd = &cobra.Command{
	Use:     "good",
	Aliases: []string{"old"},
	Short:   "mark know-good revisions",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_goodCmd).Standalone()
	bisectCmd.AddCommand(bisect_goodCmd)

	carapace.Gen(bisect_goodCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(bisect_goodCmd).DashAnyCompletion(
		carapace.ActionPositional(bisect_goodCmd),
	)
}
