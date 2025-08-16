package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles an input file into a supported output format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().StringS("format", "f", "", "The format of the output file, inferred from the extension by default")
	compileCmd.Flags().String("root", "", "Configures the project root (for absolute paths)")
	compileCmd.Flags().String("input", "", "Add a string key-value pair visible through `sys.inputs`")
	compileCmd.Flags().String("font-path", "", "Adds additional directories that are recursively searched for fonts.")
	compileCmd.Flags().Bool("ignore-system-fonts", false, "Ensures system fonts won't be searched, unless explicitly included via `--font-path`")
	compileCmd.Flags().String("package-path", "", "Custom path to local packages, defaults to system-dependent location")
	compileCmd.Flags().String("package-cache-path", "", "Custom path to package cache, defaults to system-dependent location")
	compileCmd.Flags().String("creation-timestamp", "", "The document's creation date formatted as a UNIX timestamp.")
	compileCmd.Flags().String("pages", "", "Which pages to export. When unspecified, all pages are exported.")
	compileCmd.Flags().String("pdf-standard", "", "One (or multiple comma-separated) PDF standards that Typst will enforce conformance with")
	compileCmd.Flags().Int("ppi", 144, "The PPI (pixels per inch) to use for PNG export")
	compileCmd.Flags().String("make-deps", "", "File path to which a Makefile with the current compilation's dependencies will be written")
	compileCmd.Flags().Int("jobs", 1, "Number of parallel jobs spawned during compilation. Defaults to number of CPUs. Setting it to 1 disables parallelism")
	compileCmd.Flags().String("features", "", "Enables in-development features that may be changed or removed at any time")
	compileCmd.Flags().String("diagnostic-format", "", "The format to emit diagnostics in")
	compileCmd.Flags().String("open", "", "Opens the output file with the default viewer or a specific program after compilation. Ignored if output is stdout")
	compileCmd.Flags().String("timings", "", "Produces performance timings of the compilation process.")
	compileCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")

	rootCmd.AddCommand(compileCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pdf", "png", "svg", "html").StyleF(style.ForKeyword),
		"root":   carapace.ActionDirectories(),
		// TODO: input key=value
		"font-path":          carapace.ActionDirectories(), // XXX: multiple?
		"package-path":       carapace.ActionDirectories(),
		"package-cache-path": carapace.ActionDirectories(),
		// TODO: creation timestamp
		// TODO: pages
		"pdf-standard":      carapace.ActionValues("1.7", "a-2b", "a-3b").StyleF(style.ForKeyword),
		"make-deps":         carapace.ActionFiles(),
		"features":          carapace.ActionValues("html").StyleF(style.ForKeyword),
		"diagnostic-format": carapace.ActionValues("human", "short").StyleF(style.ForKeyword),
	})

	carapace.Gen(compileCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)

}
