package cmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace/pkg/util"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:     "browse [<number> | <path> | <commit-sha>]",
	Short:   "Open repositories, issues, pull requests, and more in the browser",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(browseCmd).Standalone()

	browseCmd.Flags().BoolP("actions", "a", false, "Open repository actions")
	browseCmd.Flags().Bool("blame", false, "Open blame view for a file")
	browseCmd.Flags().StringP("branch", "b", "", "Select another branch by passing in the branch name")
	browseCmd.Flags().StringP("commit", "c", "", "Select another commit by passing in the commit SHA, default is the last commit")
	browseCmd.Flags().BoolP("no-browser", "n", false, "Print destination URL instead of opening the browser")
	browseCmd.Flags().BoolP("projects", "p", false, "Open repository projects")
	browseCmd.Flags().BoolP("releases", "r", false, "Open repository releases")
	browseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	browseCmd.Flags().BoolP("settings", "s", false, "Open repository settings")
	browseCmd.Flags().BoolP("wiki", "w", false, "Open repository wiki")
	browseCmd.Flag("commit").NoOptDefVal = " "
	rootCmd.AddCommand(browseCmd)

	carapace.Gen(browseCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(browseCmd), // TODO merge with tags
		"repo":   gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(browseCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if browseCmd.Flag("projects").Changed ||
				browseCmd.Flag("settings").Changed ||
				browseCmd.Flag("wiki").Changed {
				return carapace.ActionValues()
			}

			ref := browseCmd.Flag("branch").Value.String()
			if flag := browseCmd.Flag("commit"); flag.Changed {
				switch flag.Value.String() {
				case "", "last":
					commit, err := action.LastCommit()
					if err != nil {
						return carapace.ActionMessage(err.Error())
					}
					ref = commit.Sha

				default:
					ref = flag.Value.String()
				}
			}

			rooted := false
			if !strings.HasPrefix(c.Value, "/") && !browseCmd.Flag("repo").Changed && c.Getenv("GH_REPO") != "" {
				// GH_REPO pins the repository root, so the value is
				// relative to the current directory until it is shifted
				// to be relative to the repository root.
				root, err := util.FindReverse(c.Dir, ".git")
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				rel, err := filepath.Rel(filepath.Dir(root), c.Dir)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				abs := fmt.Sprintf("%v/%v", rel, c.Value)
				r := regexp.MustCompile(`[^\/]+\/\.\.\/`)
				for {
					if match := r.FindString(abs); match != "" {
						abs = strings.Replace(abs, match, "", 1)
					} else {
						break
					}
				}
				if rooted = !strings.Contains(abs, "/../") && !strings.HasSuffix(abs, "/.."); rooted {
					c.Value = abs
				}
			}

			contents := action.ActionContents(browseCmd, ref)
			if rooted {
				return contents.Invoke(c).ToA()
			}
			return contents
		}),
	)
}
