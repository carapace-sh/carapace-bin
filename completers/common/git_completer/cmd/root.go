package cmd

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "git",
	Short: "the stupid content tracker",
	Long:  "https://git-scm.com/",
}

const (
	group_main = iota
	group_manipulator
	group_interrogator
	group_interaction
	group_alias
	group_external
	group_low_level_manipulator
	group_low_level_interrogator
	group_low_level_synching
	group_low_level_helper
)

var groups = []*cobra.Group{
	{ID: "main", Title: "Main Commands"},
	{ID: "manipulator", Title: "Manipulator Commands"},
	{ID: "interrogator", Title: "Interrogator Commands"},
	{ID: "interaction", Title: "Interaction Commands"},
	{ID: "alias", Title: "Alias Commands"},
	{ID: "external", Title: "External Commands"},
	{ID: "low-level manipulator", Title: "Low-level Manipulator Commands"},
	{ID: "low-level interrogator", Title: "Low-level Interrogator Commands"},
	{ID: "low-level synching", Title: "Low-level Synching Commands"},
	{ID: "low-level helper", Title: "Low-level Helper Commands"},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddGroup(groups...)

	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().StringS("C", "C", "", "run as if git was started in given path")
	rootCmd.Flags().Bool("bare", false, "use $PWD as repository")
	rootCmd.Flags().StringS("c", "c", "", "pass configuration parameter to command")
	rootCmd.Flags().String("exec-path", "", "path containing core git-programs")
	rootCmd.Flags().String("git-dir", "", "path to repository")
	rootCmd.Flags().Bool("help", false, "display help message")
	rootCmd.Flags().Bool("html-path", false, "display path to HTML documentation and exit")
	rootCmd.Flags().Bool("info-path", false, "print the path where the info files are installed and exit")
	rootCmd.Flags().String("list-cmds", "", "list commands")
	rootCmd.Flags().Bool("literal-pathspecs", false, "treat pathspecs literally, rather than as glob patterns")
	rootCmd.Flags().Bool("man-path", false, "print the manpath for the man pages for this version of Git and exit")
	rootCmd.Flags().String("namespace", "", "set the Git namespace")
	rootCmd.Flags().BoolP("no-pager", "P", false, "don't pipe git output into a pager")
	rootCmd.Flags().Bool("no-replace-objects", false, "do not use replacement refs to replace git objects")
	rootCmd.Flags().BoolP("paginate", "p", false, "pipe output into a pager")
	rootCmd.Flags().Bool("version", false, "display version information")
	rootCmd.Flags().String("work-tree", "", "path to working tree")

	rootCmd.Flag("list-cmds").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
		"c": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return git.ActionConfigs().NoSpace()
			default:
				return carapace.ActionValues()
			}
		}),
		"exec-path": carapace.ActionDirectories(),
		"git-dir":   carapace.ActionDirectories(),
		"list-cmds": carapace.ActionValues("main", "others", "alias", "nohelpers"),
		"work-tree": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := rootCmd.Flag("git-dir"); f.Changed {
				gitDir, err := c.Abs(f.Value.String())
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				c.Setenv("GIT_DIR", gitDir)
			}
			if f := rootCmd.Flag("work-tree"); f.Changed {
				workTree, err := c.Abs(f.Value.String())
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				c.Setenv("GIT_WORK_TREE", workTree)
				action = action.Chdir(f.Value.String())
			}
			if f := rootCmd.Flag("C"); f.Changed {
				action = action.Chdir(f.Value.String())
			}
			return action.Invoke(c).ToA()
		})
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if cmd, _, _ := rootCmd.Find(args); cmd == rootCmd && len(args) > 0 {
			carapace.LOG.Println("adding aliases")
			addAliasCompletion(args)
			addOtherCommands()
		}
	})
}

func addAliasCompletion(args []string) {
	cmd := &cobra.Command{}
	cmd.FParseErrWhitelist.UnknownFlags = true
	cmd.Flags().StringS("C", "C", "", "run as if git was started in given path")
	cmd.Flags().String("git-dir", "", "path to repository")

	cmd.ParseFlags(args[:len(args)-1])
	aliases, err := git.Aliases(cmd.Flag("C").Value.String(), cmd.Flag("git-dir").Value.String())
	if err != nil {
		carapace.LOG.Println(err.Error())
		return
	}

	for key, value := range aliases {
		if _, _, err := rootCmd.Find([]string{key}); err == nil {
			continue // don't clobber existing commands
		}

		aliasCmd := &cobra.Command{
			Use:                key,
			Short:              fmt.Sprintf("alias for '%s'", value),
			GroupID:            groups[group_alias].ID,
			DisableFlagParsing: true,
		}

		rootCmd.AddCommand(aliasCmd)

		if strings.HasPrefix(value, "!") {
			continue // aliases beginning with ! are arbitrary shell commands so don't add completion
		}

		tokens, err := shlex.Split(value) // TODO trim value?
		if err == nil {
			carapace.Gen(aliasCmd).PositionalAnyCompletion(
				carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					c.Args = append(tokens.Strings(), c.Args...)
					return bridge.ActionCarapaceBin("git").Invoke(c).ToA()
				}),
			)
		}
	}
}

func addOtherCommands() {
	if output, err := (carapace.Context{}).Command("git", "--list-cmds=others").Output(); err == nil {
		lines := strings.Split(string(output), "\n")

		for _, name := range lines[:len(lines)-1] {
			pluginCmd := &cobra.Command{
				Use:                name,
				Short:              othersDescription(name),
				Run:                func(cmd *cobra.Command, args []string) {},
				GroupID:            groups[group_external].ID,
				DisableFlagParsing: true,
			}

			carapace.Gen(pluginCmd).PositionalAnyCompletion(
				bridge.ActionCarapaceBin("git-" + name),
			)

			rootCmd.AddCommand(pluginCmd)
		}
	}
}

func othersDescription(name string) string {
	return map[string]string{
		"abort":                    "Abort current git operation",
		"alias":                    "Define, search and show aliases",
		"archive-file":             "Export the current HEAD of the git repository to an archive",
		"authors":                  "Generate authors report",
		"browse":                   "View the web page for the current repository",
		"browse-ci":                "Opens the current git repository CI page in your default web browser",
		"brv":                      "List branches sorted by their last commit date",
		"bulk":                     "Run git commands on multiple repositories",
		"changelog":                "Generate a changelog report",
		"clang-format":             "run clang-format on lines that differ",
		"clear":                    "Rigorously clean up a repository",
		"clear-soft":               "Soft clean up a repository",
		"coauthor":                 "Add a co-author to the last commit",
		"commits-since":            "Show commit logs since some date",
		"contrib":                  "Show user's contributions",
		"count":                    "Show commit count",
		"cp":                       "Copy a file keeping its history",
		"create-branch":            "Create branches",
		"delete-branch":            "Delete branches",
		"delete-merged-branches":   "Delete merged branches",
		"delete-squashed-branches": "Delete branches that were squashed",
		"delete-submodule":         "Delete submodules",
		"delete-tag":               "Delete tags",
		"delta":                    "Lists changed files",
		"effort":                   "Show effort statistics on file(s)",
		"extras":                   "Awesome GIT utilities",
		"feature":                  "Create/Merge feature branch",
		"force-clone":              "overwrite local repositories with clone",
		"fork":                     "Fork a repo on github",
		"fresh-branch":             "Create fresh branches",
		"gh-pages":                 "Create the GitHub Pages branch",
		"graft":                    "Merge and destroy a given branch",
		"guilt":                    "calculate change between two revisions",
		"ignore":                   "Add .gitignore patterns",
		"ignore-io":                "Get sample gitignore file",
		"info":                     "Returns information on current repository",
		"local-commits":            "List local commits",
		"lock":                     "Lock a file excluded from version control",
		"locked":                   "ls files that have been locked",
		"magic":                    "Automate add/commit/push routines",
		"merge-into":               "Merge one branch into another",
		"merge-repo":               "Merge two repo histories",
		"missing":                  "Show commits missing from another branch",
		"mr":                       "Checks out a merge request locally",
		"obliterate":               "rewrite past commits to remove some files",
		"paste":                    "Send patches to pastebin for chat conversations",
		"pr":                       "Checks out a pull request locally",
		"psykorebase":              "Rebase a branch with a merge commit",
		"pull-request":             "Create pull request for GitHub project",
		"reauthor":                 "Rewrite history to change author's identity",
		"rebase-patch":             "Rebases a patch",
		"release":                  "Commit, tag and push changes to the repository",
		"rename-branch":            "rename local branch and push to remote",
		"rename-remote":            "Rename a remote",
		"rename-tag":               "Rename a tag",
		"repl":                     "git read-eval-print-loop",
		"reset-file":               "Reset one file",
		"root":                     "show path of root",
		"rscp":                     "Copies specific files from the working directory of a remote repository to the current working directory.",
		"scp":                      "Copy files to SSH compatible git-remote",
		"sed":                      "replace patterns in git-controlled files",
		"setup":                    "Set up a git repository",
		"show-merged-branches":     "Show merged branches",
		"show-tree":                "show branch tree of commit history",
		"show-unmerged-branches":   "Show unmerged branches",
		"squash":                   "squash N last changes up to a ref'ed commit",
		"stamp":                    "Stamp the last commit message",
		"standup":                  "Recall the commit history",
		"summary":                  "Show repository summary",
		"sync":                     "Sync local branch with remote branch",
		"touch":                    "Touch and add file to the index",
		"undo":                     "Remove latest commits",
		"unlock":                   "Unlock a file excluded from version control",
		"utimes":                   "Change files modification time to their last commit date",
	}[name]
}
