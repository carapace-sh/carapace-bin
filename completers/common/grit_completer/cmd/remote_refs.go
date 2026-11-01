package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_refsCmd = &cobra.Command{
	Use:   "refs",
	Short: "List references on a remote (like `git ls-remote`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_refsCmd).Standalone()

	remote_refsCmd.Flags().Bool("heads", false, "Only show refs under `refs/heads/`")
	remote_refsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	remote_refsCmd.Flags().StringP("limit", "n", "", "Return at most this many refs after sorting")
	remote_refsCmd.Flags().Bool("reverse", false, "Reverse sort order (`git ls-remote --sort=-…`)")
	remote_refsCmd.Flags().String("sort", "name", "Sort refs by `name` or `version` (`git ls-remote --sort=version:refname`)")
	remote_refsCmd.Flags().Bool("tags", false, "Only show refs under `refs/tags/`")
	remoteCmd.AddCommand(remote_refsCmd)

	carapace.Gen(remote_refsCmd).FlagCompletion(carapace.ActionMap{
		"sort": carapace.ActionValuesDescribed(
			"name", "Full ref name order",
			"version", "Numeric `version:refname` order",
		),
	})

	carapace.Gen(remote_refsCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
