package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var revParseCmd = &cobra.Command{
	Use:     "rev-parse",
	Short:   "Pick out and massage parameters",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(revParseCmd).Standalone()

	revParseCmd.Flags().String("abbrev-ref", "", "a non-ambiguous short name of the objects name")
	revParseCmd.Flags().Bool("absolute-git-dir", false, "like --git-dir, but its output is always the canonicalized absolute path")
	revParseCmd.Flags().String("after", "", "parse the date string, and output the corresponding --max-age= parameter")
	revParseCmd.Flags().Bool("all", false, "show all refs found in refs/")
	revParseCmd.Flags().String("before", "", "parse the date string, and output the corresponding --min-age= parameter")
	revParseCmd.Flags().String("branches", "", "show all branches matching pattern")
	revParseCmd.Flags().String("default", "", "if there is no parameter given by the user, use <arg> instead")
	revParseCmd.Flags().String("exclude", "", "exclude refs matching pattern")
	revParseCmd.Flags().String("exclude-hidden", "", "exclude refs hidden by given option")
	revParseCmd.Flags().Bool("flags", false, "do not output non-flag parameters")
	revParseCmd.Flags().Bool("git-common-dir", false, "show $GIT_COMMON_DIR if defined, else $GIT_DIR")
	revParseCmd.Flags().Bool("git-dir", false, "show $GIT_DIR if defined")
	revParseCmd.Flags().String("git-path", "", "resolve \"$GIT_DIR/<path>\" and takes other path relocation variables")
	revParseCmd.Flags().String("glob", "", "show all refs matching pattern")
	revParseCmd.Flags().Bool("is-inside-git-dir", false, "true, when the current working directory is below the repository directory")
	revParseCmd.Flags().Bool("is-inside-work-tree", false, "true, when the current working directory is inside the work tree of the repository")
	revParseCmd.Flags().Bool("is-shallow-repository", false, "true, when the repository is shallow")
	revParseCmd.Flags().Bool("keep-dashdash", false, "echo out the first -- met instead of skipping it")
	revParseCmd.Flags().Bool("local-env-vars", false, "list the GIT_* environment variables that are local to the repository")
	revParseCmd.Flags().Bool("no-flags", false, "do not output flag parameters")
	revParseCmd.Flags().Bool("no-revs", false, "do not output flags and parameters meant for git rev-list command")
	revParseCmd.Flags().Bool("not", false, "prefix object names with ^ and strip ^ prefix from the object names that already have one.")
	revParseCmd.Flags().String("output-object-format", "", "allow oids to be input from any object format that the current repository supports")
	revParseCmd.Flags().Bool("parseopt", false, "use in option parsing mode")
	revParseCmd.Flags().String("path-format", "", "controls the behavior of certain other options")
	revParseCmd.Flags().String("prefix", "", "behave as if git rev-parse was invoked from the <arg> subdirectory of the working tree")
	revParseCmd.Flags().BoolP("quiet", "q", false, "do not output an error message if the first argument is not a valid object name")
	revParseCmd.Flags().String("remotes", "", "show all remotes matching pattern")
	revParseCmd.Flags().String("resolve-git-dir", "", "check if <path> is a valid repository or a gitfile that points at a valid repository")
	revParseCmd.Flags().Bool("revs-only", false, "do not output flags and parameters not meant for git rev-list command")
	revParseCmd.Flags().Bool("shared-index-path", false, "show the path to the shared index file in split index mode")
	revParseCmd.Flags().String("short", "", "same as --verify but shortens the object name to a unique prefix with at least length characters")
	revParseCmd.Flags().Bool("show-cdup", false, "show the path of the top-level directory relative to the current directory")
	revParseCmd.Flags().String("show-object-format", "", "show the object format used for the repository for storage")
	revParseCmd.Flags().Bool("show-prefix", false, "show the path of the current directory relative to the top-level directory")
	revParseCmd.Flags().Bool("show-ref-format", false, "show the reference storage format used for the repository")
	revParseCmd.Flags().Bool("show-superproject-working-tree", false, "show the absolute path of the root of the superproject’s working tree")
	revParseCmd.Flags().Bool("show-toplevel", false, "show the path of the top-level directory of the working tree")
	revParseCmd.Flags().String("since", "", "parse the date string, and output the corresponding --max-age= parameter")
	revParseCmd.Flags().Bool("sq", false, "output a single line")
	revParseCmd.Flags().Bool("sq-quote", false, "use in shell quoting mode")
	revParseCmd.Flags().Bool("stop-at-non-option", false, "stop at the first non-option argument")
	revParseCmd.Flags().Bool("stuck-long", false, "output the options in their long form")
	revParseCmd.Flags().Bool("symbolic", false, "output object names in a form as close to the original input as possible")
	revParseCmd.Flags().Bool("symbolic-full-name", false, "similar to --symbolic, but omits input that are not refs")
	revParseCmd.Flags().String("tags", "", "show all tags matching pattern")
	revParseCmd.Flags().String("until", "", "parse the date string, and output the corresponding --min-age= parameter")
	revParseCmd.Flags().Bool("verify", false, "verify that exactly one parameter is provided")
	rootCmd.AddCommand(revParseCmd)

	revParseCmd.Flag("short").NoOptDefVal = " "
	revParseCmd.Flag("abbrev-ref").NoOptDefVal = " "
	revParseCmd.Flag("branches").NoOptDefVal = " "
	revParseCmd.Flag("show-object-format").NoOptDefVal = " "

	carapace.Gen(revParseCmd).FlagCompletion(carapace.ActionMap{
		"abbrev-ref":           carapace.ActionValues("strict", "loose").StyleF(style.ForKeyword),
		"after":                time.ActionDate(),
		"before":               time.ActionDate(),
		"branches":             carapace.ActionValues(),
		"exclude":              carapace.ActionValues(),
		"exclude-hidden":       carapace.ActionValues("fetch", "receive", "uploadpack"),
		"git-path":             carapace.ActionFiles().ChdirF(traverse.GitDir),
		"glob":                 carapace.ActionValues(),
		"output-object-format": carapace.ActionValues("sha1", "sha256", "storage"),
		"path-format":          carapace.ActionValues("absolute", "relative").StyleF(style.ForKeyword),
		"prefix":               carapace.ActionDirectories().ChdirF(traverse.GitWorkTree),
		"remotes":              carapace.ActionValues(),
		"resolve-git-dir":      carapace.ActionFiles(),
		"show-object-format":   carapace.ActionValues("storage", "input", "output"),
		"since":                time.ActionDate(),
		"tags":                 carapace.ActionValues(),
		"until":                time.ActionDate(),
	})
}
