package cmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	_git "github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action/git"
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
					commit, err := _git.LastCommit()
					if err != nil {
						return carapace.ActionMessage(err.Error())
					}
					ref = commit.Sha

				default:
					ref = flag.Value.String()
				}
			}

			path := filepath.Dir(c.Value)
			path = strings.TrimPrefix(path, "/")
			if !strings.HasPrefix(c.Value, "/") && !browseCmd.Flag("repo").Changed && c.Getenv("GH_REPO") != "" {
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
				path = filepath.Dir(abs)
			}
			// TODO merge with issue/pr ids (show all with prefix search using callbackvalue)
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return action.ActionContents(browseCmd, path, ref)
			})
		}),
	)
}
