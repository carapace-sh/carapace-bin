package common

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func AddFontFlags(cmd *cobra.Command) {
	cmd.Flags().String("font-path", "", "Adds additional directories that are recursively searched for fonts.")
	cmd.Flags().Bool("ignore-embedded-fonts", false, "Ensures fonts embedded into Typst won't be considered")
	cmd.Flags().Bool("ignore-system-fonts", false, "Ensures system fonts won't be searched, unless explicitly included via `--font-path`")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"font-path": carapace.ActionDirectories().List(string(os.PathListSeparator)),
	})
}

func AddPackageFlags(cmd *cobra.Command) {
	cmd.Flags().String("package-cache-path", "", "Custom path to package cache, defaults to system-dependent location")
	cmd.Flags().String("package-path", "", "Custom path to local packages, defaults to system-dependent location")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"package-cache-path": carapace.ActionDirectories(),
		"package-path":       carapace.ActionDirectories(),
	})
}

func AddBuildFlags(cmd *cobra.Command) {
	cmd.Flags().Uint("creation-timestamp", 0, "The document's creation date formatted as a UNIX timestamp.")
	cmd.Flags().String("diagnostic-format", "", "The format to emit diagnostics in")
	cmd.Flags().String("features", "", "Enables in-development features that may be changed or removed at any time")
	cmd.Flags().StringP("format", "f", "", "The format of the output file, inferred from the extension by default")
	cmd.Flags().String("input", "", "Add a string key-value pair visible through `sys.inputs`")
	cmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs spawned during compilation. Defaults to number of CPUs. Setting it to 1 disables parallelism")
	cmd.Flags().String("root", "", "Configures the project root (for absolute paths)")
	AddFontFlags(cmd)
	AddPackageFlags(cmd)

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"diagnostic-format": carapace.ActionValues("human", "short"),
		"features":          carapace.ActionValues("html", "a11y-extras").StyleF(style.ForExtension),
		"format":            carapace.ActionValues("pdf", "png", "svg", "html").StyleF(style.ForExtension),
		"root":              carapace.ActionDirectories(),
	})
}

func AddPdfFlags(cmd *cobra.Command) {
	cmd.Flags().String("deps", "", "File path to which a list of current compilation's dependencies will be written. Use `-` to write to stdout")
	cmd.Flags().String("deps-format", "", "File format to use for dependencies")
	cmd.Flags().String("make-deps", "", "File path to which a Makefile with the current compilation's dependencies will be written")
	cmd.Flags().Bool("no-pdf-tags", false, "By default, even when not producing a `PDF/UA-1` document, a tagged PDF document is written to provide a baseline of accessibility. In some circumstances (for example when trying to reduce the size of a document) it can be desirable to disable tagged PDF")
	cmd.Flags().String("open", "", "Opens the output file with the default viewer or a specific program after compilation. Ignored if output is stdout")
	cmd.Flags().String("pages", "", "Which pages to export. When unspecified, all pages are exported.")
	cmd.Flags().String("pdf-standard", "", "One (or multiple comma-separated) PDF standards that Typst will enforce conformance with")
	cmd.Flags().Int("ppi", 144, "The PPI (pixels per inch) to use for PNG export")
	cmd.Flags().String("timings", "", "Produces performance timings of the compilation process.")
	AddBuildFlags(cmd)

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		// TODO: pages
		"deps": carapace.ActionFiles(),
		"deps-format": carapace.ActionValuesDescribed(
			"json", "Encodes as JSON, failing for non-Unicode paths",
			"zero", "Separates paths with NULL bytes and can express all paths",
			"make", "Emits in Make format, omitting inexpressible paths",
		),
		"make-deps": carapace.ActionFiles(),
		"pdf-standard": carapace.ActionValuesDescribed(
			"1.4", "PDF 1.4",
			"1.5", "PDF 1.5",
			"1.6", "PDF 1.6",
			"1.7", "PDF 1.7",
			"2.0", "PDF 2.0",
			"a-1b", "PDF/A-1b",
			"a-1a", "PDF/A-1a",
			"a-2b", "PDF/A-2b",
			"a-2u", "PDF/A-2u",
			"a-2a", "PDF/A-2a",
			"a-3b", "PDF/A-3b",
			"a-3u", "PDF/A-3u",
			"a-3a", "PDF/A-3a",
			"a-4", "PDF/A-4",
			"a-4f", "PDF/A-4f",
			"a-4e", "PDF/A-4e",
			"ua-1", "PDF/UA-1",
		),
	})
}
