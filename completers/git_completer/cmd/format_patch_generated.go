package cmd

import (
	"github.com/spf13/cobra"
)

var format_patchCmd = &cobra.Command{
	Use:   "format-patch",
	Short: "Prepare patches for e-mail submission",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	format_patchCmd.Flags().String("add-header", "", "add email header")
	format_patchCmd.Flags().String("attach", "", "attach the patch")
	format_patchCmd.Flags().String("base", "", "add prerequisite tree info to the patch series")
	format_patchCmd.Flags().String("cc", "", "add Cc: header")
	format_patchCmd.Flags().String("cover-from-description", "", "generate parts of a cover letter based on a branch's description")
	format_patchCmd.Flags().Bool("cover-letter", false, "generate a cover letter")
	format_patchCmd.Flags().String("creation-factor", "", "percentage by which creation is weighted")
	format_patchCmd.Flags().String("from", "", "set From address to <ident> (or committer ident if absent)")
	format_patchCmd.Flags().Bool("ignore-if-in-upstream", false, "don't include a patch matching a commit upstream")
	format_patchCmd.Flags().String("in-reply-to", "", "make first mail a reply to <message-id>")
	format_patchCmd.Flags().String("inline", "", "inline the patch")
	format_patchCmd.Flags().String("interdiff", "", "show changes against <rev> in cover letter or single patch")
	format_patchCmd.Flags().BoolP("keep-subject", "k", false, "don't strip/add [PATCH]")
	format_patchCmd.Flags().Bool("no-binary", false, "don't output binary diffs")
	format_patchCmd.Flags().BoolP("no-numbered", "N", false, "use [PATCH] even with multiple patches")
	format_patchCmd.Flags().BoolP("no-stat", "p", false, "show patch format instead of default (patch + stat)")
	format_patchCmd.Flags().BoolP("numbered", "n", false, "use [PATCH n/m] even with a single patch")
	format_patchCmd.Flags().Bool("numbered-files", false, "use simple number sequence for output file names")
	format_patchCmd.Flags().BoolP("output-directory", "o", false, "<dir>    store resulting files in <dir>")
	format_patchCmd.Flags().Bool("progress", false, "show progress while generating patches")
	format_patchCmd.Flags().BoolP("quiet", "q", false, "don't print the patch filenames")
	format_patchCmd.Flags().String("range-diff", "", "show changes against <refspec> in cover letter or single patch")
	format_patchCmd.Flags().BoolP("reroll-count", "v", false, "<n>    mark the series as Nth re-roll")
	format_patchCmd.Flags().Bool("rfc", false, "Use [RFC PATCH] instead of [PATCH]")
	format_patchCmd.Flags().String("signature", "", "add a signature")
	format_patchCmd.Flags().String("signature-file", "", "add a signature from a file")
	format_patchCmd.Flags().BoolP("signoff", "s", false, "add Signed-off-by:")
	format_patchCmd.Flags().String("start-number", "", "start numbering patches at <n> instead of 1")
	format_patchCmd.Flags().Bool("stdout", false, "print patches to standard out")
	format_patchCmd.Flags().String("subject-prefix", "", "Use [<prefix>] instead of [PATCH]")
	format_patchCmd.Flags().String("suffix", "", "use <sfx> instead of '.patch'")
	format_patchCmd.Flags().String("thread", "", "enable message threading, styles: shallow, deep")
	format_patchCmd.Flags().String("to", "", "add To: header")
	format_patchCmd.Flags().Bool("zero-commit", false, "output all-zero hash in From header")
	rootCmd.AddCommand(format_patchCmd)
}
