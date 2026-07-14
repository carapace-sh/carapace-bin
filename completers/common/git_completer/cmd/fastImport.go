package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fastImportCmd = &cobra.Command{
	Use:     "fast-import",
	Short:   "Backend for fast Git data importers",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {
	carapace.Gen(fastImportCmd).Standalone()

	fastImportCmd.Flags().String("active-branches", "", "maximum number of branches to maintain active at once")
	fastImportCmd.Flags().Bool("allow-unsafe-features", false, "enable unsafe options")
	fastImportCmd.Flags().String("big-file-threshold", "", "maximum size of a blob that fast-import will attempt to create a delta for")
	fastImportCmd.Flags().String("cat-blob-fd", "", "write responses to get-mark, cat-blob, and ls queries to the file descriptor <fd> instead of stdout")
	fastImportCmd.Flags().String("date-format", "", "specify the type of dates the frontend will supply")
	fastImportCmd.Flags().String("depth", "", "maximum delta depth, for blob and tree deltification")
	fastImportCmd.Flags().Bool("done", false, "terminate with error if there is no done command at the end of the stream")
	fastImportCmd.Flags().String("export-marks", "", "dumps the internal marks table to <file> when complete")
	fastImportCmd.Flags().String("export-pack-edges", "", "after creating a packfile, print a line of data to <file>")
	fastImportCmd.Flags().Bool("force", false, "force updating modified existing branches")
	fastImportCmd.Flags().String("import-marks", "", "before processing any input, load the marks specified in <file>")
	fastImportCmd.Flags().String("import-marks-if-exists", "", "like --import-marks but instead of erroring out")
	fastImportCmd.Flags().String("max-pack-size", "", "maximum size of each output packfile")
	fastImportCmd.Flags().Bool("no-relative-marks", false, "paths are not relative to an internal direcotyr")
	fastImportCmd.Flags().Bool("quiet", false, "disable the output shown by --stats")
	fastImportCmd.Flags().Bool("relative-marks", false, "paths are relative to an internal directory")
	fastImportCmd.Flags().String("rewrite-submodules-from", "", "rewrite the object IDs for the submodule specified by <name>")
	fastImportCmd.Flags().String("rewrite-submodules-to", "", "rewrite the object IDs for the submodule specified by <name>")
	fastImportCmd.Flags().String("signed-commits", "", "handle signed commits according to mode")
	fastImportCmd.Flags().String("signed-tags", "", "handle signed tags according to mode")
	fastImportCmd.Flags().Bool("stats", false, "display some basic statistics about the objects fast-import has created")
	rootCmd.AddCommand(fastImportCmd)

	carapace.Gen(fastImportCmd).FlagCompletion(carapace.ActionMap{
		"cat-blob-fd":             carapace.ActionFiles(),
		"export-marks":            carapace.ActionFiles(), // TODO support relative
		"export-pack-edges":       carapace.ActionValues(),
		"import-marks":            carapace.ActionFiles(), // TODO support relative
		"import-marks-if-exists":  carapace.ActionFiles(),
		"rewrite-submodules-from": carapace.ActionValues(), // TODO
		"rewrite-submodules-to":   carapace.ActionValues(), // TODO
		"signed-commits": carapace.ActionValuesDescribed(
			"verbatim", "store signatures as-is",
			"warn-verbatim", "warn if signature does not verify and store as-is",
			"warn-strip", "warn if signature does not verify and strip it",
			"strip", "strip any signature",
			"strip-if-invalid", "strip only invalid signatures",
			"sign-if-invalid", "replace invalid signatures with newly created ones",
			"abort", "abort if signature does not verify",
		),
		"signed-tags": carapace.ActionValuesDescribed(
			"verbatim", "store signatures as-is",
			"warn-verbatim", "warn if signature does not verify and store as-is",
			"warn-strip", "warn if signature does not verify and strip it",
			"strip", "strip any signature",
			"strip-if-invalid", "strip only invalid signatures",
			"sign-if-invalid", "replace invalid signatures with newly created ones",
			"abort", "abort if signature does not verify",
		),
	})
}
