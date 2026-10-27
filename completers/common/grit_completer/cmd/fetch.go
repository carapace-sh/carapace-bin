package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Download refs and objects from a remote",
	GroupID: "remote",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().Bool("all", false, "Fetch from every configured remote (respects `remote.<name>.skipFetchAll`)")
	fetchCmd.Flags().String("deepen", "", "Deepen an existing shallow clone by this many commits (`git fetch --deepen`)")
	fetchCmd.Flags().String("depth", "", "Fetch only the recent history of each branch (same as `git fetch --depth`)")
	fetchCmd.Flags().Bool("dry-run", false, "Report ref updates without writing refs or objects")
	fetchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	fetchCmd.Flags().Bool("no-prune", false, "Do not prune stale remote-tracking refs, even when config enables it")
	fetchCmd.Flags().Bool("no-tags", false, "Do not fetch tags automatically (`git fetch --no-tags`)")
	fetchCmd.Flags().Bool("prune", false, "Remove remote-tracking refs that no longer exist on the remote")
	fetchCmd.Flags().BoolP("quiet", "q", false, "Suppress progress and human output")
	fetchCmd.Flags().Bool("tags", false, "Fetch all tags from the remote (`git fetch --tags`)")
	fetchCmd.Flags().Bool("unshallow", false, "Convert a shallow repository to a complete one (`git fetch --unshallow`)")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
