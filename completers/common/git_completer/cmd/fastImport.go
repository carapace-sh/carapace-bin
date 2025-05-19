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

	fastImportCmd.Flags().String("active-branches", "", "Maximum number of branches to maintain active at once")
	fastImportCmd.Flags().Bool("allow-unsafe-features", false, "Enable unsafe options")
	fastImportCmd.Flags().String("big-file-threshold", "", "Maximum size of a blob that fast-import will attempt to create a delta for")
	fastImportCmd.Flags().String("cat-blob-fd", "", "Write responses to get-mark, cat-blob, and ls queries to the file descriptor <fd> instead of stdout")
	fastImportCmd.Flags().String("date-format", "", "Specify the type of dates the frontend will supply")
	fastImportCmd.Flags().String("depth", "", "Maximum delta depth, for blob and tree deltification")
	fastImportCmd.Flags().Bool("done", false, "Terminate with error if there is no done command at the end of the stream")
	fastImportCmd.Flags().String("export-marks", "", "Dumps the internal marks table to <file> when complete")
	fastImportCmd.Flags().String("export-pack-edges", "", "After creating a packfile, print a line of data to <file>")
	fastImportCmd.Flags().Bool("force", false, "Force updating modified existing branches")
	fastImportCmd.Flags().String("import-marks", "", "Before processing any input, load the marks specified in <file>")
	fastImportCmd.Flags().String("import-marks-if-exists", "", "Like --import-marks but instead of erroring out")
	fastImportCmd.Flags().String("max-pack-size", "", "Maximum size of each output packfile")
	fastImportCmd.Flags().Bool("no-relative-marks", false, "Paths are not relative to an internal direcotyr")
	fastImportCmd.Flags().Bool("quiet", false, "Disable the output shown by --stats")
	fastImportCmd.Flags().Bool("relative-marks", false, "Paths are relative to an internal directory")
	fastImportCmd.Flags().String("rewrite-submodules-from", "", "Rewrite the object IDs for the submodule specified by <name>")
	fastImportCmd.Flags().String("rewrite-submodules-to", "", "Rewrite the object IDs for the submodule specified by <name>")
	fastImportCmd.Flags().Bool("stats", false, "Display some basic statistics about the objects fast-import has created")
	rootCmd.AddCommand(fastImportCmd)

	carapace.Gen(fastImportCmd).FlagCompletion(carapace.ActionMap{
		"cat-blob-fd":             carapace.ActionFiles(),
		"export-marks":            carapace.ActionFiles(), // TODO support relative
		"export-pack-edges":       carapace.ActionValues(),
		"import-marks":            carapace.ActionFiles(), // TODO support relative
		"import-marks-if-exists":  carapace.ActionFiles(),
		"rewrite-submodules-from": carapace.ActionValues(), // TODO
		"rewrite-submodules-to":   carapace.ActionValues(), // TODO
	})
}
