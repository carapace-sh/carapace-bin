package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Download objects and refs from another repository",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().Bool("all", false, "fetch from all remotes")
	fetchCmd.Flags().BoolP("append", "a", false, "append to .git/FETCH_HEAD instead of overwriting")
	fetchCmd.Flags().Bool("atomic", false, "Use an atomic transaction to update local refs")
	fetchCmd.Flags().Bool("auto-gc", false, "run git maintenance run --auto at the end")
	fetchCmd.Flags().Bool("auto-maintenance", false, "run git maintenance run --auto at the end")
	fetchCmd.Flags().String("deepen", "", "deepen history of shallow clone")
	fetchCmd.Flags().String("depth", "", "deepen history of shallow clone")
	fetchCmd.Flags().Bool("dry-run", false, "dry run")
	fetchCmd.Flags().StringArray("filter", nil, "object filtering")
	fetchCmd.Flags().BoolP("force", "f", false, "force overwrite of local reference")
	fetchCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	fetchCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	fetchCmd.Flags().StringP("jobs", "j", "", "number of submodules fetched in parallel")
	fetchCmd.Flags().BoolP("keep", "k", false, "keep downloaded pack")
	fetchCmd.Flags().BoolP("multiple", "m", false, "fetch from multiple remotes")
	fetchCmd.Flags().Bool("negotiate-only", false, "do not fetch anything from the server")
	fetchCmd.Flags().String("negotiation-tip", "", "report that we have only objects reachable from this object")
	fetchCmd.Flags().Bool("no-auto-gc", false, "do not run git maintenance run --auto at the end")
	fetchCmd.Flags().Bool("no-auto-maintenance", false, "do not run git maintenance run --auto at the end")
	fetchCmd.Flags().Bool("no-recurse-submodules", false, "disable recursive fetching of submodules")
	fetchCmd.Flags().BoolP("no-tags", "n", false, "disable automatic tag following")
	fetchCmd.Flags().Bool("no-write-fetch-head", false, "do not write the list of remote refs fetched")
	fetchCmd.Flags().Bool("prefetch", false, "place all refs into the refs/prefetch/ namespace")
	fetchCmd.Flags().Bool("progress", false, "force progress reporting")
	fetchCmd.Flags().BoolP("prune", "p", false, "prune remote-tracking branches no longer on remote")
	fetchCmd.Flags().BoolP("prune-tags", "P", false, "prune local tags no longer on remote and clobber changed tags")
	fetchCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	fetchCmd.Flags().String("recurse-submodules", "", "control recursive fetching of submodules")
	fetchCmd.Flags().String("recurse-submodules-default", "", "used internally to temporarily provide a non-negative default value")
	fetchCmd.Flags().String("refmap", "", "specify fetch refmap")
	fetchCmd.Flags().StringP("server-option", "o", "", "option to transmit")
	fetchCmd.Flags().Bool("set-upstream", false, "set upstream for git pull/fetch")
	fetchCmd.Flags().String("shallow-exclude", "", "deepen history of shallow clone, excluding rev")
	fetchCmd.Flags().String("shallow-since", "", "deepen history of shallow repository based on time")
	fetchCmd.Flags().Bool("show-forced-updates", false, "check for forced-updates on all updated branches")
	fetchCmd.Flags().Bool("stdin", false, "Read refspecs, one per line, from stdin")
	fetchCmd.Flags().String("submodule-prefix", "", "prepend <path> to paths printed in informative messages")
	fetchCmd.Flags().BoolP("tags", "t", false, "fetch all tags and associated objects")
	fetchCmd.Flags().Bool("unshallow", false, "convert to a complete repository")
	fetchCmd.Flags().BoolP("update-head-ok", "u", false, "allow updating of HEAD ref")
	fetchCmd.Flags().Bool("update-shallow", false, "accept refs that update .git/shallow")
	fetchCmd.Flags().String("upload-pack", "", "path to upload pack on remote end")
	fetchCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	fetchCmd.Flags().Bool("write-commit-graph", false, "write the commit-graph after fetching")
	fetchCmd.Flags().Bool("write-fetch-head", false, "write the list of remote refs fetched")

	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"filter":                     git.ActionObjectFilters(),
		"recurse-submodules":         carapace.ActionValues("yes", "on-demand").StyleF(style.ForKeyword),
		"recurse-submodules-default": carapace.ActionValues("yes", "on-demand").StyleF(style.ForKeyword),
	})

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// from `man git-fetch`:
			// 1. git fetch [<options>] [<repository> [<refspec>...]]
			// 2. git fetch [<options>] <group>
			// 3. git fetch --multiple [<options>] [(<repository> | <group>)...]
			// 4. git fetch --all [<options>]

			// 4. git fetch --all [<options>]
			// no positional completion
			if fetchCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}

			// 3. git fetch --multiple [<options>] [(<repository> | <group>)...]
			// <repository> and <group> completion
			if fetchCmd.Flag("multiple").Changed {
				return git.ActionRemotes().FilterArgs()
			}

			// 1. git fetch [<options>] [<repository> [<refspec>...]]
			// <repository> and <refspec> completion
			//
			// 2. git fetch [<options>] <group>
			// <group> completion

			// if we have no remote (<repository> or <group>) yet, complete that
			if len(c.Args) == 0 {
				return git.ActionRemotes()
			}

			remote := c.Args[0]
			return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
				if len(c.Parts) == 0 {
					return git.ActionRemoteBranchNames(remote)
				}

				return git.ActionLocalBranches()
			})
		}),
	)
}
