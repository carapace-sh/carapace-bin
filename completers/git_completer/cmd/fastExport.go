package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var fastExportCmd = &cobra.Command{
	Use:     "fast-export",
	Short:   "Git data exporter",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {
	carapace.Gen(fastExportCmd).Standalone()

	fastExportCmd.Flags().BoolS("C", "C", false, "Perform copy detection")
	fastExportCmd.Flags().BoolS("M", "M", false, "Perform move detection")
	fastExportCmd.Flags().Bool("anonymize", false, "Anonymize the contents of the repository while still retaining the shape of the history and stored tree")
	fastExportCmd.Flags().StringSlice("anonymize-map", nil, "Convert token <from> to <to> in the anonymized output")
	fastExportCmd.Flags().String("export-marks", "", "Dumps the internal marks table to <file> when complete")
	fastExportCmd.Flags().Bool("export-marks.", false, "")
	fastExportCmd.Flags().Bool("fake-missing-tagger", false, "Fake a tagger to be able to fast-import the output")
	fastExportCmd.Flags().Bool("full-tree", false, "This option will cause fast-export to issue a \"deleteall\" directive for each commit")
	fastExportCmd.Flags().String("import-marks", "", "Before processing any input, load the marks specified in <file>")
	fastExportCmd.Flags().Bool("mark-tags", false, "In addition to labelling blobs and commits with mark ids, also label tags")
	fastExportCmd.Flags().Bool("no-data", false, "Skip output of blob objects and instead refer to blobs via their original SHA-1 hash")
	fastExportCmd.Flags().String("progress", "", "Insert progress statements every <n> objects")
	fastExportCmd.Flags().String("reencode", "", "Specify how to handle encoding header in commit objects")
	fastExportCmd.Flags().Bool("reference-excluded-parents", false, "Refer to commits in the excluded range of history by their sha1sum")
	fastExportCmd.Flags().Bool("refspec", false, "Apply the specified refspec to each ref exported")
	fastExportCmd.Flags().Bool("show-original-ids", false, "Add an extra directive to the output for commits and blobs, original-oid <SHA1SUM>")
	fastExportCmd.Flags().String("signed-tags", "", "Specify how to handle signed tags")
	fastExportCmd.Flags().String("tag-of-filtered-object", "", "Specify how to handle tags whose tagged object is filtered out")
	fastExportCmd.Flags().Bool("use-done-feature", false, "Start the stream with a feature done stanza, and terminate it with a done command")
	rootCmd.AddCommand(fastExportCmd)

	carapace.Gen(fastExportCmd).FlagCompletion(carapace.ActionMap{
		"anonymize-map":          carapace.ActionValues(), // TODO
		"export-marks":           carapace.ActionFiles(),
		"import-marks":           carapace.ActionFiles(),
		"progress":               carapace.ActionValues(),
		"reencode":               carapace.ActionValues("yes", "no", "abort").StyleF(style.ForKeyword),
		"signed-tags":            carapace.ActionValues("verbatim", "warn", "warn-strip", "strip", "abort"),
		"tag-of-filtered-object": carapace.ActionValues("abort", "drop", "rewrite"),
	})

	carapace.Gen(fastExportCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("git", "rev-list"),
	)
}
