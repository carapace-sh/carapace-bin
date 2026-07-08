package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge another branch into the current one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{
			LocalBranches:  true,
			RemoteBranches: true,
			Tags:           true,
		}).FilterArgs(),
	)

	carapace.Gen(mergeCmd).DashAnyCompletion(
		carapace.ActionPositional(mergeCmd),
	)
}
