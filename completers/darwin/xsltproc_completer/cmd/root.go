package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xsltproc",
	Short: "command line XSLT processor",
	Long:  "https://man.freebsd.org/cgi/man.cgi?xsltproc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("catalogs", false, "Use the SGML catalog specified in SGML_CATALOG_FILES")
	rootCmd.Flags().Bool("debug", false, "Output an XML tree of the transformed document")
	rootCmd.Flags().Bool("dumpextensions", false, "Dump the list of all registered extensions")
	rootCmd.Flags().String("encoding", "", "Specify the output encoding")
	rootCmd.Flags().Bool("html", false, "The input document is an HTML file")
	rootCmd.Flags().Bool("huge", false, "Relax hardcoded limits of the XML parser")
	rootCmd.Flags().Bool("load-trace", false, "Display all the documents loaded during processing")
	rootCmd.Flags().String("maxdepth", "", "Adjust the maximum depth of the template stack")
	rootCmd.Flags().String("maxparserdepth", "", "Maximum element nesting level")
	rootCmd.Flags().String("maxvars", "", "Maximum number of variables")
	rootCmd.Flags().Bool("nodtdattr", false, "Do not apply default attributes from the document's DTD")
	rootCmd.Flags().Bool("nomkdir", false, "Do not create directories")
	rootCmd.Flags().Bool("nonet", false, "Do not use the Internet to fetch DTDs or entities")
	rootCmd.Flags().Bool("noout", false, "Suppress output")
	rootCmd.Flags().Bool("novalid", false, "Do not validate")
	rootCmd.Flags().Bool("nowrite", false, "Do not write to files")
	rootCmd.Flags().String("output", "", "Specify output file or directory")
	rootCmd.Flags().String("param", "", "Pass a parameter to the stylesheet")
	rootCmd.Flags().String("path", "", "Use the list of filesystem paths to load DTDs or entities")
	rootCmd.Flags().Bool("profile", false, "Output profiling information")
	rootCmd.Flags().Bool("norman", false, "Do not profile")
	rootCmd.Flags().Bool("repeat", false, "Repeat 100 times for timing")
	rootCmd.Flags().String("seed-rand", "", "Initialize pseudo random number generator")
	rootCmd.Flags().String("stringparam", "", "Pass a string parameter to the stylesheet")
	rootCmd.Flags().Bool("timing", false, "Output timing information")
	rootCmd.Flags().Bool("verbose", false, "Verbose")
	rootCmd.Flags().Bool("version", false, "Version")
	rootCmd.Flags().Bool("xinclude", false, "Do XInclude processing")
	rootCmd.Flags().Bool("xincludestyle", false, "Do XInclude processing on stylesheet")
	rootCmd.Flags().String("writesubtree", "", "Restrict file writing to a subtree")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"encoding":     carapace.ActionValues("UTF-8", "UTF-16", "ISO-8859-1", "ASCII"),
		"output":       carapace.ActionFiles(),
		"path":         carapace.ActionDirectories(),
		"writesubtree": carapace.ActionDirectories(),
	})
}
