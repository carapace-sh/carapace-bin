package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

func AddDiffFlags(cmd *cobra.Command) {
	cmd.Flags().BoolS("0", "0", false, "omit diff output for unmerged entries")
	cmd.Flags().StringS("O", "O", "", "control the order in which files appear in the output")
	cmd.Flags().BoolS("R", "R", false, "swap two inputs")
	cmd.Flags().StringS("S", "S", "", "look for differences that change the number of occurrences of the specified string")
	cmd.Flags().String("abbrev", "", "show only a partial prefix")
	cmd.Flags().StringArray("anchored", nil, "generate a diff using the \"anchored diff\" algorithm")
	cmd.Flags().BoolP("base", "1", false, "compare with base")
	cmd.Flags().Bool("binary", false, "output a binary diff")
	cmd.Flags().StringP("break-rewrites", "B", "", "break complete rewrite changes into pairs of delete and create")
	cmd.Flags().Bool("cached", false, "view the changes you staged in the index/cache")
	cmd.Flags().BoolP("cc", "c", false, "combined diff format for merge commits")
	cmd.Flags().Bool("check", false, "warn if changes introduce conflict markers or whitespace errors")
	cmd.Flags().String("color", "", "show colored diff")
	cmd.Flags().String("color-moved", "", "moved lines of code are colored differently")
	cmd.Flags().String("color-moved-ws", "", "configure how whitespace is ignored when performing move detection")
	cmd.Flags().String("color-words", "", "equivalent to --word-diff=color plus (if a regex was specified)")
	cmd.Flags().Bool("compact-summary", false, "output a condensed summary of extended header information")
	cmd.Flags().Bool("cumulative", false, "synonym for --dirstat=cumulative")
	cmd.Flags().Bool("default-prefix", false, "use the default source and destination prefixes")
	cmd.Flags().String("diff-algorithm", "", "choose a diff algorithm")
	cmd.Flags().String("diff-filter", "", "filter files")
	cmd.Flags().StringP("dirstat", "X", "", "output the distribution of relative amount of changes for each sub-directory")
	cmd.Flags().String("dirstat-by-file", "", "synonym for --dirstat=files,param1,param2...")
	cmd.Flags().String("dst-prefix", "", "show the given destination prefix instead of \"b/\"")
	cmd.Flags().Bool("exit-code", false, "make the program exit with codes similar to diff(1)")
	cmd.Flags().Bool("ext-diff", false, "allow an external diff helper to be executed")
	cmd.Flags().StringP("find-copies", "C", "", "detect copies as well as renames")
	cmd.Flags().Bool("find-copies-harder", false, "inspect unmodified files as candidates for the source of copy")
	cmd.Flags().String("find-object", "", "look for differences that change the number of occurrences of the specified object")
	cmd.Flags().StringP("find-renames", "M", "", "detect renames")
	cmd.Flags().String("follow", "", "continue listing the history of a file beyond renames")
	cmd.Flags().Bool("full-index", false, "show the full pre- and post-image blob object names")
	cmd.Flags().BoolP("function-context", "W", false, "show whole surrounding functions of changes")
	cmd.Flags().Bool("histogram", false, "generate a diff using the \"histogram diff\" algorithm")
	cmd.Flags().BoolP("ignore-all-space", "w", false, "ignore whitespace when comparing lines")
	cmd.Flags().Bool("ignore-blank-lines", false, "ignore changes whose lines are all blank")
	cmd.Flags().Bool("ignore-cr-at-eol", false, "ignore carriage-return at the end of line when doing a comparison")
	cmd.Flags().StringArrayP("ignore-matching-lines", "I", nil, "ignore changes whose all lines match <regex>")
	cmd.Flags().Bool("ignore-space-at-eol", false, "ignore changes in whitespace at EOL")
	cmd.Flags().BoolP("ignore-space-change", "b", false, "ignore changes in amount of whitespace")
	cmd.Flags().String("ignore-submodules", "", "ignore changes to submodules in the diff generation")
	cmd.Flags().Bool("indent-heuristic", false, "enable the heuristic that shifts diff hunk boundaries to make patches easier to read")
	cmd.Flags().String("inter-hunk-context", "", "show the context between diff hunks")
	cmd.Flags().BoolP("irreversible-delete", "D", false, "omit the preimage for deletes")
	cmd.Flags().Bool("ita-invisible-in-index", false, "this option makes the entry appear as a new file")
	cmd.Flags().StringS("l", "l", "", "prevent rename/copy detection from running if the number of rename/copy targets exceeds the specified number")
	cmd.Flags().String("line-prefix", "", "prepend an additional prefix to every line of output")
	cmd.Flags().Bool("minimal", false, "spend extra time to make sure the smallest possible diff is produced")
	cmd.Flags().Bool("name-only", false, "show only names of changed files")
	cmd.Flags().Bool("name-status", false, "show only names and status of changed files")
	cmd.Flags().Bool("no-abbrev", false, "show the full 40-byte hexadecimal commit object name")
	cmd.Flags().Bool("no-color", false, "turn off colored diff")
	cmd.Flags().Bool("no-color-moved", false, "turn off move detection")
	cmd.Flags().Bool("no-color-moved-ws", false, "do not ignore whitespace when performing move detection")
	cmd.Flags().Bool("no-ext-diff", false, "disallow external diff drivers")
	cmd.Flags().Bool("no-follow", false, "do not continue listing the history of a file beyond renames")
	cmd.Flags().Bool("no-indent-heuristic", false, "disable the indent heuristic")
	cmd.Flags().Bool("no-index", false, "compare paths on the file system")
	cmd.Flags().BoolP("no-patch", "s", false, "suppress diff output")
	cmd.Flags().Bool("no-prefix", false, "do not show any source or destination prefix")
	cmd.Flags().Bool("no-rename-empty", false, "whether to use empty blobs as rename source")
	cmd.Flags().Bool("no-renames", false, "turn off rename detection")
	cmd.Flags().String("no-textconv", "", "disallow external text conversion filters to be run when comparing binary files")
	cmd.Flags().Bool("numstat", false, "similar to --stat, but shows number of added and deleted lines in decimal notation")
	cmd.Flags().BoolP("ours", "2", false, "compare with our branch")
	cmd.Flags().String("output", "", "output to a specific file instead of stdout")
	cmd.Flags().String("output-indicator-context", "", "specify the character used to indicate context lines in the generated patch")
	cmd.Flags().String("output-indicator-new", "", "specify the character used to indicate new lines in the generated patch")
	cmd.Flags().String("output-indicator-old", "", "specify the character used to indicate old lines in the generated patch")
	cmd.Flags().BoolP("patch", "p", false, "generate patch")
	cmd.Flags().Bool("patch-with-raw", false, "synonym for -p --raw")
	cmd.Flags().Bool("patch-with-stat", false, "synonym for -p --stat")
	cmd.Flags().Bool("patience", false, "generate a diff using the \"patience diff\" algorithm")
	cmd.Flags().Bool("pickaxe-all", false, "when -S or -G finds a change, show all the changes in that changeset")
	cmd.Flags().Bool("pickaxe-regex", false, "treat the <string> given to -S as an extended POSIX regular expression to match")
	cmd.Flags().Bool("raw", false, "generate the diff in raw format")
	cmd.Flags().String("relative", "", "exclude changes outside the directory")
	cmd.Flags().Bool("rename-empty", false, "whether to use empty blobs as rename source")
	cmd.Flags().String("rotate-to", "", "move the files before the named <file> to the end")
	cmd.Flags().Bool("shortstat", false, "output only the last line of the --stat format")
	cmd.Flags().String("skip-to", "", "discard the files before the named <file> from the output")
	cmd.Flags().String("src-prefix", "", "show the given source prefix instead of \"a/\"")
	cmd.Flags().String("stat", "", "generate a diffstat")
	cmd.Flags().String("stat-count", "", "generate diffstat with limited lines")
	cmd.Flags().String("stat-graph-width", "", "generate diffstat with a given graph width")
	cmd.Flags().String("stat-name-width", "", "generate diffstat with a given filename width")
	cmd.Flags().String("stat-width", "", "generate diffstat with a given width")
	cmd.Flags().String("submodule", "", "specify how differences in submodules are shown")
	cmd.Flags().Bool("summary", false, "output a condensed summary of extended header information")
	cmd.Flags().BoolP("text", "a", false, "treat all files as text")
	cmd.Flags().String("textconv", "", "allow external text conversion filters to be run when comparing binary files")
	cmd.Flags().BoolP("theirs", "3", false, "compare with their branch")
	cmd.Flags().BoolS("u", "u", false, "generate patch")
	cmd.Flags().StringP("unified", "U", "", "generate diffs with <n> lines of context instead of the usual three")
	cmd.Flags().String("word-diff", "", "show a word diff, using the <mode> to delimit changed words")
	cmd.Flags().String("word-diff-regex", "", "use <regex> to decide what a word is")
	cmd.Flags().String("ws-error-highlight", "", "highlight whitespace errors in the context, old or new lines of the diff")
	cmd.Flags().BoolS("z", "z", false, "do not munge pathnames and use NULs as output field terminators")

	cmd.Flag("color").NoOptDefVal = " "
	cmd.Flag("color-moved").NoOptDefVal = "default"
	cmd.Flag("color-moved-ws").NoOptDefVal = " "
	cmd.Flag("word-diff").NoOptDefVal = "plain"

	cmd.MarkFlagsMutuallyExclusive("abbrev", "no-abbrev")
	cmd.MarkFlagsMutuallyExclusive("color", "no-color")
	cmd.MarkFlagsMutuallyExclusive("color-moved", "no-color-moved")
	cmd.MarkFlagsMutuallyExclusive("color-moved-ws", "no-color-moved-ws")
	cmd.MarkFlagsMutuallyExclusive("ext-diff", "no-ext-diff")
	cmd.MarkFlagsMutuallyExclusive("follow", "no-follow")
	cmd.MarkFlagsMutuallyExclusive("indent-heuristic", "no-indent-heuristic")
	cmd.MarkFlagsMutuallyExclusive("patch", "no-patch")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"color":              git.ActionColorModes(),
		"color-moved":        git.ActionColorMovedModes(),
		"color-moved-ws":     git.ActionColorMovedWsModes(),
		"diff-algorithm":     git.ActionDiffAlgorithms(),
		"follow":             carapace.ActionFiles(), // TODO complete files of specific revision/modified between commits?
		"ignore-submodules":  carapace.ActionValues("none", "untracked", "dirty", "all"),
		"output":             carapace.ActionFiles(),
		"rotate-to":          carapace.ActionFiles(), // TODO complete files of specific revision/modified between commits?
		"skip-to":            carapace.ActionFiles(), // TODO complete files of specific revision/modified between commits?
		"submodule":          carapace.ActionValues("short", "long", "log"),
		"word-diff":          git.ActionWordDiffModes(),
		"ws-error-highlight": git.ActionWsErrorHighlightModes().UniqueList(","),
	})
}

// AddPatchContextFlags adds flags for patch-mode context control
// used by commands like add, checkout, commit, reset, restore, stash push.
func AddPatchContextFlags(cmd *cobra.Command) {
	cmd.Flags().String("inter-hunk-context", "", "show context between diff hunks up to the specified number of lines")
	cmd.Flags().StringP("unified", "U", "", "generate diffs with <n> lines of context instead of the usual three")
}
