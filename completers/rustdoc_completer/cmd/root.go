package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rustdoc"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rustdoc",
	Long:  "https://doc.rust-lang.org/1.71.0/rustdoc/what-is-rustdoc.html",
	Short: "generate documentation for Rust projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("Z", "Z", "", "unstable / perma-unstable options (only on nightly build)")
	rootCmd.Flags().StringP("allow", "A", "", "Set lint allowed")
	rootCmd.Flags().String("cap-lints", "", "Set the most restrictive lint level")
	rootCmd.Flags().Bool("cfg", false, "pass a --cfg to rustc")
	rootCmd.Flags().Bool("check", false, "Run rustdoc checks")
	rootCmd.Flags().Bool("check-cfg", false, "pass a --check-cfg to rustc")
	rootCmd.Flags().String("check-theme", "", "check if given theme is valid")
	rootCmd.Flags().StringP("codegen", "C", "", "pass a codegen option to rustc")
	rootCmd.Flags().String("color", "", "Configure coloring of output")
	rootCmd.Flags().String("crate-name", "", "specify the name of this crate")
	rootCmd.Flags().String("crate-type", "", "Comma separated list of types of crates for the compiler to emit")
	rootCmd.Flags().String("crate-version", "", "crate version to print into documentation")
	rootCmd.Flags().String("default-setting", "", "Default value for a rustdoc setting")
	rootCmd.Flags().String("default-theme", "", "Set the default theme")
	rootCmd.Flags().StringP("deny", "D", "", "Set lint denied")
	rootCmd.Flags().String("diagnostic-width", "", "Provide width of the output for truncated error messages")
	rootCmd.Flags().Bool("disable-minification", false, "removed")
	rootCmd.Flags().Bool("disable-per-crate-search", false, "disables generating the crate selector on the search box")
	rootCmd.Flags().Bool("display-doctest-warnings", false, "show warnings that originate in doctests")
	rootCmd.Flags().Bool("document-hidden-items", false, "document items that have doc(hidden)")
	rootCmd.Flags().Bool("document-private-items", false, "document private items")
	rootCmd.Flags().String("edition", "", "edition to use when compiling rust code")
	rootCmd.Flags().String("emit", "", "Comma separated list of types of output for rustdoc to emit")
	rootCmd.Flags().Bool("enable-index-page", false, "To enable generation of the index page")
	rootCmd.Flags().Bool("enable-per-target-ignores", false, "parse ignore-foo for ignoring doctests on a per-target basis")
	rootCmd.Flags().String("error-format", "", "How errors and other messages are produced")
	rootCmd.Flags().StringP("extend-css", "e", "", "To add some CSS rules with a given file to generate doc with your own theme")
	rootCmd.Flags().String("extern", "", "pass an --extern to rustc")
	rootCmd.Flags().Bool("extern-html-root-takes-precedence", false, "give precedence to `--extern-html-root-url`")
	rootCmd.Flags().String("extern-html-root-url", "", "base URL to use for dependencies")
	rootCmd.Flags().StringP("forbid", "F", "", "Set lint forbidden")
	rootCmd.Flags().String("force-warn", "", "Set lint force-warn")
	rootCmd.Flags().Bool("generate-link-to-definition", false, "Make the identifiers in the HTML source code pages navigable")
	rootCmd.Flags().Bool("generate-redirect-map", false, "Generate JSON file at the top level instead of generating HTML redirection files")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message")
	rootCmd.Flags().String("html-after-content", "", "files to include inline between the content and </body> of a rendered Markdown file")
	rootCmd.Flags().String("html-before-content", "", "files to include inline between <body> and the content of a rendered Markdown file")
	rootCmd.Flags().String("html-in-header", "", "files to include inline in the <head> section of a rendered Markdown file")
	rootCmd.Flags().String("index-page", "", "Markdown file to be used as index page")
	rootCmd.Flags().String("json", "", "Configure the structure of JSON diagnostics")
	rootCmd.Flags().StringP("library-path", "L", "", "directory to add to crate search path")
	rootCmd.Flags().String("markdown-after-content", "", "files to include inline between the content and </body> of a rendered Markdown file")
	rootCmd.Flags().String("markdown-before-content", "", "files to include inline between <body> and the content of a rendered Markdown file")
	rootCmd.Flags().String("markdown-css", "", "CSS files to include via <link> in a rendered Markdown file")
	rootCmd.Flags().Bool("markdown-no-toc", false, "don't include table of contents")
	rootCmd.Flags().String("markdown-playground-url", "", "URL to send code snippets to")
	rootCmd.Flags().Bool("no-run", false, "Compile doctests without running them")
	rootCmd.Flags().Bool("nocapture", false, "Don't capture stdout and stderr of tests")
	rootCmd.Flags().StringP("out-dir", "o", "", "which directory to place the output")
	rootCmd.Flags().String("output", "", "Which directory to place the output")
	rootCmd.Flags().StringP("output-format", "w", "", "the output type to write")
	rootCmd.Flags().String("persist-doctests", "", "Directory to persist doctest executables into")
	rootCmd.Flags().String("playground-url", "", "URL to send code snippets to")
	rootCmd.Flags().String("resource-suffix", "", "suffix to add to CSS and JavaScript files")
	rootCmd.Flags().Bool("runtool", false, "The tool to run tests with when building for a different target than host")
	rootCmd.Flags().Bool("runtool-arg", false, "One (of possibly many) arguments to pass to the runtool")
	rootCmd.Flags().Bool("scrape-examples-output-path", false, "collect function call information and output at the given path")
	rootCmd.Flags().Bool("scrape-examples-target-crate", false, "collect function call information for functions from the target crate")
	rootCmd.Flags().Bool("scrape-tests", false, "Include test code when scraping examples")
	rootCmd.Flags().Bool("show-coverage", false, "calculate percentage of public items with documentation")
	rootCmd.Flags().Bool("show-type-layout", false, "Include the memory layout of types in the docs")
	rootCmd.Flags().Bool("sort-modules-by-appearance", false, "sort modules by where they appear in the program, rather than alphabetically")
	rootCmd.Flags().String("static-root-path", "", "Path string to force loading static files from in output pages")
	rootCmd.Flags().String("sysroot", "", "Override the system root")
	rootCmd.Flags().String("target", "", "target triple to document")
	rootCmd.Flags().Bool("test", false, "run code examples as tests")
	rootCmd.Flags().String("test-args", "", "arguments to pass to the test runner")
	rootCmd.Flags().String("test-builder", "", "The rustc-like binary to use as the test builder")
	rootCmd.Flags().String("test-run-directory", "", "The working directory in which to run tests")
	rootCmd.Flags().String("theme", "", "additional themes which will be added to the generated docs")
	rootCmd.Flags().BoolP("verbose", "v", false, "use verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "print rustdoc's version")
	rootCmd.Flags().StringP("warn", "W", "", "Set lint warnings")
	rootCmd.Flags().Bool("with-examples", false, "path to function call information")

	// TODO flag completion
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// "Z":                       carapace.ActionValues(),
		// "allow":                   carapace.ActionValues(),
		// "cap-lints":               carapace.ActionValues(),
		"check-theme": carapace.ActionFiles().List(","),
		// "codegen":                 carapace.ActionValues(),
		"color":      carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"crate-name": carapace.ActionValues(),
		"crate-type": carapace.ActionValues("bin", "lib", "rlib", "dylib", "cdylib", "staticlib", "proc-macro"),
		// "crate-version":           carapace.ActionValues(),
		// "default-setting":         carapace.ActionValues(),
		"default-theme": carapace.ActionFiles().List(","),
		"deny":          rustdoc.ActionLints(),
		// "diagnostic-width":        carapace.ActionValues(),
		// "edition":                 carapace.ActionValues(),
		"emit":         carapace.ActionValues("unversioned-shared-resources", "toolchain-shared-resources", "invocation-specific"),
		"error-format": carapace.ActionValues("human", "json", "short"),
		// "extend-css":              carapace.ActionValues(),
		// "extern":                  carapace.ActionValues(),
		// "extern-html-root-url":    carapace.ActionValues(),
		"forbid":              rustdoc.ActionLints(),
		"force-warn":          rustdoc.ActionLints(),
		"html-after-content":  carapace.ActionFiles().List(","),
		"html-before-content": carapace.ActionFiles().List(","),
		"html-in-header":      carapace.ActionFiles().List(","),
		// "json":                    carapace.ActionValues(),
		"library-path":            carapace.ActionDirectories(),
		"markdown-after-content":  carapace.ActionFiles().List(","),
		"markdown-before-content": carapace.ActionFiles().List(","),
		"markdown-css":            carapace.ActionFiles().List(","),
		// "markdown-playground-url": carapace.ActionValues(),
		"out-dir":          carapace.ActionDirectories(),
		"output-format":    carapace.ActionValues("html"),
		"persist-doctests": carapace.ActionDirectories(),
		// "playground-url":     carapace.ActionValues(),
		// "resource-suffix":    carapace.ActionValues(),
		// "static-root-path":   carapace.ActionValues(),
		"sysroot": carapace.ActionDirectories(),
		// "target":             carapace.ActionValues(),
		// "test-args":          carapace.ActionValues(),
		// "test-builder":       carapace.ActionValues(),
		"test-run-directory": carapace.ActionDirectories(),
		// "theme":              carapace.ActionValues(),
		"warn": rustdoc.ActionLints(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles().Prefix("@"),
	)
}
