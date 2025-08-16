package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"os"
)

func AddPackageFlags(cmd *cobra.Command) {
	cmd.Flags().String("package-path", "", "Custom path to local packages, defaults to system-dependent location")
	cmd.Flags().String("package-cache-path", "", "Custom path to package cache, defaults to system-dependent location")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"package-path":       carapace.ActionDirectories(),
		"package-cache-path": carapace.ActionDirectories(),
	})
}

func AddBuildFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("format", "f", "", "The format of the output file, inferred from the extension by default")
	cmd.Flags().String("root", "", "Configures the project root (for absolute paths)")
	cmd.Flags().String("input", "", "Add a string key-value pair visible through `sys.inputs`")
	cmd.Flags().String("font-path", "", "Adds additional directories that are recursively searched for fonts.")
	cmd.Flags().Bool("ignore-system-fonts", false, "Ensures system fonts won't be searched, unless explicitly included via `--font-path`")
	cmd.Flags().Uint("creation-timestamp", 0, "The document's creation date formatted as a UNIX timestamp.")
	cmd.Flags().Int("jobs", 1, "Number of parallel jobs spawned during compilation. Defaults to number of CPUs. Setting it to 1 disables parallelism")
	cmd.Flags().String("features", "", "Enables in-development features that may be changed or removed at any time")
	cmd.Flags().String("diagnostic-format", "", "The format to emit diagnostics in")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"format":            carapace.ActionValues("pdf", "png", "svg", "html").StyleF(style.ForKeyword),
		"root":              carapace.ActionDirectories(),
		"font-path":         carapace.ActionDirectories().List(string(os.PathListSeparator)),
		"features":          carapace.ActionValues("html").StyleF(style.ForKeyword),
		"diagnostic-format": carapace.ActionValues("human", "short").StyleF(style.ForKeyword),
	})

	AddPackageFlags(cmd)
}

func AddPdfFlags(cmd *cobra.Command) {
	cmd.Flags().String("pages", "", "Which pages to export. When unspecified, all pages are exported.")
	cmd.Flags().String("pdf-standard", "", "One (or multiple comma-separated) PDF standards that Typst will enforce conformance with")
	cmd.Flags().Int("ppi", 144, "The PPI (pixels per inch) to use for PNG export")
	cmd.Flags().String("make-deps", "", "File path to which a Makefile with the current compilation's dependencies will be written")
	cmd.Flags().String("open", "", "Opens the output file with the default viewer or a specific program after compilation. Ignored if output is stdout")
	cmd.Flags().String("timings", "", "Produces performance timings of the compilation process.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		// TODO: pages
		"pdf-standard": carapace.ActionValues("1.7", "a-2b", "a-3b").StyleF(style.ForKeyword),
		"make-deps":    carapace.ActionFiles(),
	})

	AddBuildFlags(cmd)
}
