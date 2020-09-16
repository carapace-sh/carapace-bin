package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/git_completer/cmd/action"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show changes between commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().StringS("l", "l", "", "prevent rename/copy detection from running if the number of rename/copy targets exceeds the specified number")
	diffCmd.Flags().BoolS("u", "u", false, "Generate patch")
	diffCmd.Flags().BoolS("z", "z", false, "do not munge pathnames and use NULs as output field terminators")
	diffCmd.Flags().StringS("O", "O", "", "Control the order in which files appear in the output.")
	diffCmd.Flags().BoolS("0", "0", false, "Omit diff output for unmerged entries")
	diffCmd.Flags().BoolS("R", "R", false, "Swap two inputs")
	diffCmd.Flags().StringS("S", "S", "", "Look for differences that change the number of occurrences of the specified string")
	diffCmd.Flags().String("abbrev", "", "show only a partial prefix")
	diffCmd.Flags().String("anchored", "", "Generate a diff using the \"anchored diff\" algorithm.")
	diffCmd.Flags().BoolP("base", "1", false, "compare with base")
	diffCmd.Flags().Bool("binary", false, "output a binary diff")
	diffCmd.Flags().StringP("break-rewrites", "B", "", "Break complete rewrite changes into pairs of delete and create.")
	diffCmd.Flags().Bool("check", false, "Warn if changes introduce conflict markers or whitespace errors")
	diffCmd.Flags().String("color", "", "Show colored diff.")
	diffCmd.Flags().String("color-moved", "", "Moved lines of code are colored differently.")
	diffCmd.Flags().String("color-moved-ws", "", "This configures how whitespace is ignored when performing the move detection")
	diffCmd.Flags().String("color-words", "", "Equivalent to --word-diff=color plus (if a regex was specified)")
	diffCmd.Flags().Bool("compact-summary", false, "Output a condensed summary of extended header information")
	diffCmd.Flags().Bool("cumulative", false, "Synonym for --dirstat=cumulative")
	diffCmd.Flags().String("diff-algorithm", "", "")
	diffCmd.Flags().String("diff-filter", "", "filter files")
	diffCmd.Flags().StringP("dirstat", "X", "", "Output the distribution of relative amount of changes for each sub-directory")
	diffCmd.Flags().String("dirstat-by-file", "", "Synonym for --dirstat=files,param1,param2...")
	diffCmd.Flags().String("dst-prefix", "", "Show the given destination prefix instead of \"b/\".")
	diffCmd.Flags().Bool("exit-code", false, "Make the program exit with codes similar to diff(1).")
	diffCmd.Flags().Bool("ext-diff", false, "Allow an external diff helper to be executed.")
	diffCmd.Flags().StringP("find-copies", "C", "", "Detect copies as well as renames")
	diffCmd.Flags().Bool("find-copies-harder", false, "inspect unmodified files as candidates for the source of copy")
	diffCmd.Flags().String("find-object", "", "Look for differences that change the number of occurrences of the specified object.")
	diffCmd.Flags().StringP("find-renames", "M", "", "Detect renames")
	diffCmd.Flags().Bool("full-index", false, "show the full pre- and post-image blob object names")
	diffCmd.Flags().BoolP("function-context", "W", false, "Show whole surrounding functions of changes.")
	diffCmd.Flags().Bool("histogram", false, "Generate a diff using the \"histogram diff\" algorithm.")
	diffCmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines.")
	diffCmd.Flags().Bool("ignore-blank-lines", false, "Ignore changes whose lines are all blank.")
	diffCmd.Flags().Bool("ignore-cr-at-eol", false, "Ignore carriage-return at the end of line when doing a comparison.")
	diffCmd.Flags().Bool("ignore-space-at-eol", false, "Ignore changes in whitespace at EOL.")
	diffCmd.Flags().BoolP("ignore-space-change", "b", false, "Ignore changes in amount of whitespace.")
	diffCmd.Flags().String("ignore-submodules", "", "Ignore changes to submodules in the diff generation.")
	diffCmd.Flags().Bool("indent-heuristic", false, "Enable the heuristic that shifts diff hunk boundaries to make patches easier to read.")
	diffCmd.Flags().String("inter-hunk-context", "", "Show the context between diff hunks")
	diffCmd.Flags().BoolP("irreversible-delete", "D", false, "Omit the preimage for deletes")
	diffCmd.Flags().Bool("ita-invisible-in-index", false, "this option makes the entry appear as a new file")
	diffCmd.Flags().String("line-prefix", "", "Prepend an additional prefix to every line of output.")
	diffCmd.Flags().Bool("minimal", false, "Spend extra time to make sure the smallest possible diff is produced.")
	diffCmd.Flags().Bool("name-only", false, "Show only names of changed files.")
	diffCmd.Flags().Bool("name-status", false, "Show only names and status of changed files.")
	diffCmd.Flags().Bool("no-color", false, "Turn off colored diff.")
	diffCmd.Flags().Bool("no-color-moved", false, "Turn off move detection.")
	diffCmd.Flags().Bool("no-color-moved-ws", false, "Do not ignore whitespace when performing move detection")
	diffCmd.Flags().Bool("no-ext-diff", false, "Disallow external diff drivers.")
	diffCmd.Flags().Bool("no-indent-heuristic", false, "Disable the indent heuristic.")
	diffCmd.Flags().BoolP("no-patch", "s", false, "Suppress diff output.")
	diffCmd.Flags().Bool("no-prefix", false, "Do not show any source or destination prefix.")
	diffCmd.Flags().Bool("no-rename-empty", false, "Whether to use empty blobs as rename source.")
	diffCmd.Flags().Bool("no-renames", false, "Turn off rename detection")
	diffCmd.Flags().Bool("numstat", false, "Similar to --stat, but shows number of added and deleted lines in decimal notation")
	diffCmd.Flags().BoolP("ours", "2", false, "compare with our branch")
	diffCmd.Flags().String("output", "", "Output to a specific file instead of stdout.")
	diffCmd.Flags().String("output-indicator-context", "", "Specify the character used to indicate context lines in the generated patch.")
	diffCmd.Flags().String("output-indicator-new", "", "Specify the character used to indicate new lines in the generated patch.")
	diffCmd.Flags().String("output-indicator-old", "", "Specify the character used to indicate old lines in the generated patch.")
	diffCmd.Flags().BoolP("patch", "p", false, "Generate patch")
	diffCmd.Flags().Bool("patch-with-raw", false, "Synonym for -p --raw.")
	diffCmd.Flags().Bool("patch-with-stat", false, "Synonym for -p --stat.")
	diffCmd.Flags().Bool("patience", false, "Generate a diff using the \"patience diff\" algorithm.")
	diffCmd.Flags().Bool("pickaxe-all", false, "When -S or -G finds a change, show all the changes in that changeset")
	diffCmd.Flags().Bool("pickaxe-regex", false, "Treat the <string> given to -S as an extended POSIX regular expression to match.")
	diffCmd.Flags().Bool("quiet", false, "Disable all output of the program.")
	diffCmd.Flags().Bool("raw", false, "Generate the diff in raw format.")
	diffCmd.Flags().String("relative,", "", "exclude changes outside the directory")
	diffCmd.Flags().Bool("rename-empty", false, "Whether to use empty blobs as rename source.")
	diffCmd.Flags().Bool("shortstat", false, "Output only the last line of the --stat format")
	diffCmd.Flags().String("src-prefix", "", "Show the given source prefix instead of \"a/\".")
	diffCmd.Flags().String("stat", "", "")
	diffCmd.Flags().String("submodule", "", "Specify how differences in submodules are shown.")
	diffCmd.Flags().Bool("summary", false, "Output a condensed summary of extended header information")
	diffCmd.Flags().BoolP("text", "a", false, "Treat all files as text.")
	diffCmd.Flags().String("textconv,", "", "Allow (or disallow) external text conversion filters to be run when comparing binary files")
	diffCmd.Flags().BoolP("theirs", "3", false, "compare with their branch")
	diffCmd.Flags().StringP("unified", "U", "", "Generate diffs with <n> lines of context instead of the usual three.")
	diffCmd.Flags().String("word-diff", "", "Show a word diff, using the <mode> to delimit changed words.")
	diffCmd.Flags().String("word-diff-regex", "", "Use <regex> to decide what a word is")
	diffCmd.Flags().String("ws-error-highlight", "", "Highlight whitespace errors in the context, old or new lines of the diff")
	rootCmd.AddCommand(diffCmd)

	diffCmd.Flag("color-moved").NoOptDefVal = "default"
	diffCmd.Flag("color-moved-ws").NoOptDefVal = " "
	diffCmd.Flag("word-diff").NoOptDefVal = "plain"

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"color":          carapace.ActionValues("always", "never", "auto"),
		"color-moved":    ActionColorMovedModes(),
		"color-moved-ws": ActionColorMovedWsModes(),
		"diff-algorithm": ActionDiffAlgorithms(),
		"output":         carapace.ActionFiles(""),
		"submodule":      carapace.ActionValues("short", "long", "log"),
		"word-diff":      ActionWordDiffModes(),
	})

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			for _, arg := range os.Args {
				if arg == "--" {
					return carapace.ActionFiles("")
				}
			}
			return action.ActionRefs(action.RefOptionDefault)
		}),
	)
}

func ActionDiffAlgorithms() carapace.Action {
	return carapace.ActionValuesDescribed(
		"myers", "The basic greedy diff algorithm",
		"minimal", "Spend extra time to make sure the smallest possible diff is produced",
		"patience",
		"Use patience diff algorithm when generating patches",
		"histogram", "This algorithm extends the patience algorithm to support low-occurrence common elements",
	)
}

func ActionColorMovedModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "Moved lines are not highlighted",
		"default", "default mode",
		"plain", "plain mode",
		"blocks", "greedily detects blocks",
		"zebra", "Blocks of moved text are detected as in blocks mode",
		"dimmed-zebra", "Similar to zebra, but additional dimming of uninteresting parts of moved code",
	)
}

func ActionColorMovedWsModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "Do not ignore whitespace when performing move detection",
		"ignore-space-at-eol", "Ignore changes in whitespace at EOL",
		"ignore-space-change", "Ignore changes in amount of whitespace.",
		"ignore-all-space", "Ignore whitespace when comparing lines.",
		"allow-indentation-change", "Initially ignore any whitespace in the move detection",
	)
}

func ActionWordDiffModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"color", "Highlight changed words using only colors",
		"plain", "Show words as [-removed-] and {+added+}",
		"porcelain", "Use a special line-based format intended for script consumption",
		"none", "Disable word diff again",
	)
}
