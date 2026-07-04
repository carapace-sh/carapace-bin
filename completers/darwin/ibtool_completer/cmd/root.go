package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ibtool",
	Short: "Interface Builder compile tool",
	Long:  "https://keith.github.io/xcode-manpages/ibtool.1.html",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "include all information in plist output")
	rootCmd.Flags().String("bundle", "", "load the bundle at the specified path")
	rootCmd.Flags().String("cocoatouch-compiler-mode", "", "UIKit compilation mode: toolchain or simulator")
	rootCmd.Flags().String("companion-strings-file", "", "specify locale:stringsFile for a compiled document")
	rootCmd.Flags().String("compile", "", "compile the input file and write to path")
	rootCmd.Flags().String("connections", "", "include the document's connections")
	rootCmd.Flags().String("convert", "", "rename class old to new")
	rootCmd.Flags().Bool("enable-auto-layout", false, "enable Auto Layout in the document")
	rootCmd.Flags().Bool("errors", false, "include document errors in plist output")
	rootCmd.Flags().String("export", "", "export specified properties for each object")
	rootCmd.Flags().String("export-strings-file", "", "extract localizable strings into a strings file")
	rootCmd.Flags().String("export-xliff", "", "extract localizable strings into an XLIFF document")
	rootCmd.Flags().String("flatten", "", "when combined with --compile, set to NO for runnable and editable")
	rootCmd.Flags().String("hierarchy", "", "include the document's hierarchy")
	rootCmd.Flags().String("import", "", "apply property values from a plist to objects with matching IDs")
	rootCmd.Flags().String("import-strings-file", "", "replace document's localizable strings with translations from file")
	rootCmd.Flags().String("import-xliff", "", "replace document's localizable strings with translations from XLIFF")
	rootCmd.Flags().String("incremental-file", "", "specify the translated document for incremental localization")
	rootCmd.Flags().Bool("localizable-all", false, "include all localizable attributes")
	rootCmd.Flags().Bool("localizable-geometry", false, "include localizable geometry")
	rootCmd.Flags().Bool("localizable-other", false, "include localizable attributes (non-string, non-geometry)")
	rootCmd.Flags().Bool("localizable-stringarrays", false, "include localizable to-many relationship strings")
	rootCmd.Flags().Bool("localizable-strings", false, "include non-empty localizable strings")
	rootCmd.Flags().Bool("localizable-to-many-relationships", false, "include localizable to-many relationships")
	rootCmd.Flags().Bool("localize-incremental", false, "consolidate structural and localization changes")
	rootCmd.Flags().String("module", "", "specify the module name for Swift custom classes")
	rootCmd.Flags().Bool("notices", false, "include document notices in plist output")
	rootCmd.Flags().Bool("objects", false, "include the document's objects keyed by object ID")
	rootCmd.Flags().String("output-format", "", "output format: xml1, binary1, or human-readable-text")
	rootCmd.Flags().String("plugin", "", "load the plug-in located at pluginPath")
	rootCmd.Flags().String("plugin-dir", "", "load all valid plug-ins found in the first level of pluginPath")
	rootCmd.Flags().String("previous-file", "", "specify the Interface Builder document from the previous iteration")
	rootCmd.Flags().Bool("reference-external-strings-file", false, "make Base.lproj files reference external strings files")
	rootCmd.Flags().Bool("remove-plugin-dependencies", false, "remove dependencies on IB 3 plug-ins")
	rootCmd.Flags().String("source-language", "", "source language for XLIFF export")
	rootCmd.Flags().String("strip", "", "remove design-time content from a NIB, writing to path")
	rootCmd.Flags().String("target-language", "", "target language for XLIFF export")
	rootCmd.Flags().Bool("update-constraints", false, "adjust constraint constants of misplaced objects")
	rootCmd.Flags().Bool("update-frames", false, "adjust frames of misplaced/ambiguous objects")
	rootCmd.Flags().Bool("upgrade", false, "upgrade the document to the latest document type")
	rootCmd.Flags().Bool("version", false, "print version of ibtool")
	rootCmd.Flags().Bool("version-history", false, "print the IB/System Version used when document was last saved")
	rootCmd.Flags().Bool("warnings", false, "include document warnings in plist output")
	rootCmd.Flags().String("write", "", "write the resulting Interface Builder document to path")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bundle":                   carapace.ActionDirectories(),
		"cocoatouch-compiler-mode": carapace.ActionValues("toolchain", "simulator"),
		"compile":                  carapace.ActionDirectories(),
		"export-strings-file":      carapace.ActionFiles(),
		"export-xliff":             carapace.ActionFiles(),
		"flatten":                  carapace.ActionValues("YES", "NO"),
		"import":                   carapace.ActionFiles(),
		"import-strings-file":      carapace.ActionFiles(),
		"import-xliff":             carapace.ActionFiles(),
		"incremental-file":         carapace.ActionFiles(),
		"output-format":            carapace.ActionValues("xml1", "binary1", "human-readable-text"),
		"plugin":                   carapace.ActionFiles(),
		"plugin-dir":               carapace.ActionDirectories(),
		"previous-file":            carapace.ActionFiles(),
		"source-language":          carapace.ActionValues("en", "fr", "de", "es", "ja", "zh-Hans", "zh-Hant"),
		"strip":                    carapace.ActionDirectories(),
		"target-language":          carapace.ActionValues("en", "fr", "de", "es", "ja", "zh-Hans", "zh-Hant"),
		"write":                    carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
