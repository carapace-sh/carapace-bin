package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var fetchPackCmd = &cobra.Command{
	Use:     "fetch-pack",
	Short:   "Receive missing objects from another repository",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(fetchPackCmd).Standalone()

	fetchPackCmd.Flags().Bool("all", false, "fetch all remote refs")
	fetchPackCmd.Flags().Bool("check-self-contained-and-connected", false, "output \"connectivity-ok\" if the received pack is self-contained and connected")
	fetchPackCmd.Flags().Bool("deepen-relative", false, "argument --depth specifies the number of commits from the current shallow boundary instead of from the tip of each remote branch history")
	fetchPackCmd.Flags().String("depth", "", "limit fetching to ancestor-chains not longer than n")
	fetchPackCmd.Flags().String("exec", "", "same as --upload-pack=<git-upload-pack>")
	fetchPackCmd.Flags().Bool("include-tag", false, "if the remote side supports it, annotated tags objects will be downloaded on the same connection")
	fetchPackCmd.Flags().BoolP("keep", "k", false, "do not invoke git unpack-objects on received data, but create a single packfile out of it instead")
	fetchPackCmd.Flags().Bool("no-progress", false, "do not show the progress")
	fetchPackCmd.Flags().BoolP("quiet", "q", false, "pass -q flag to git unpack-objects; this makes the cloning process less verbose")
	fetchPackCmd.Flags().Bool("refetch", false, "skips negotiating commits with the server in order to fetch all matching objects")
	fetchPackCmd.Flags().String("shallow-exclude", "", "deepen or shorten the history of a shallow repository to exclude commits reachable from a specified remote branch or tag")
	fetchPackCmd.Flags().String("shallow-since", "", "deepen or shorten the history of a shallow repository to include all reachable commits after <date>")
	fetchPackCmd.Flags().Bool("stdin", false, "take the list of refs from stdin, one per line")
	fetchPackCmd.Flags().Bool("thin", false, "fetch a \"thin\" pack, which records objects in deltified form based on objects not included in the pack")
	fetchPackCmd.Flags().String("upload-pack", "", "use this to specify the path to git-upload-pack on the remote side, if is not found on your $PATH")
	fetchPackCmd.Flags().BoolS("v", "v", false, "run verbosely")
	rootCmd.AddCommand(fetchPackCmd)

	carapace.Gen(fetchPackCmd).FlagCompletion(carapace.ActionMap{
		"shallow-exclude": git.ActionRefs(git.RefOption{}.Default()),
		"shallow-since":   time.ActionDate(),
	})

	carapace.Gen(fetchPackCmd).PositionalCompletion(
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	)

	carapace.Gen(fetchPackCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{
				Url:      c.Args[0],
				Branches: true,
				Tags:     true,
			})
		}),
	)
}
