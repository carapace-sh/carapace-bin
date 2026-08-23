package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xmllint",
	Short: "command line XML tool",
	Long:  "https://man.freebsd.org/cgi/man.cgi?xmllint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("auto", false, "Generate a small document for testing purposes")
	rootCmd.Flags().Bool("catalogs", false, "Use the SGML catalog(s) from SGML_CATALOG_FILES")
	rootCmd.Flags().Bool("c14n", false, "Use W3C XML Canonicalisation")
	rootCmd.Flags().Bool("chkregister", false, "Turn on node registration")
	rootCmd.Flags().Bool("compress", false, "Turn on gzip compression of output")
	rootCmd.Flags().Bool("copy", false, "Test the internal copy implementation")
	rootCmd.Flags().Bool("debug", false, "Parse a file and output an annotated tree")
	rootCmd.Flags().Bool("debugent", false, "Debug the entities defined in the document")
	rootCmd.Flags().Bool("dropdtd", false, "Remove DTD from output")
	rootCmd.Flags().Bool("dtdattr", false, "Fetch external DTD and populate the tree with inherited attributes")
	rootCmd.Flags().String("dtdvalid", "", "Use the DTD specified by a URL for validation")
	rootCmd.Flags().String("dtdvalidfpi", "", "Use the DTD specified by a Formal Public Identifier")
	rootCmd.Flags().String("encode", "", "Output in the given encoding")
	rootCmd.Flags().Bool("format", false, "Reformat and reindent the output")
	rootCmd.Flags().Bool("help", false, "Print usage summary")
	rootCmd.Flags().Bool("html", false, "Use the HTML parser")
	rootCmd.Flags().Bool("htmlout", false, "Output results as an HTML file")
	rootCmd.Flags().Bool("insert", false, "Test for valid insertions")
	rootCmd.Flags().Bool("loaddtd", false, "Fetch an external DTD")
	rootCmd.Flags().Bool("load-trace", false, "Display all the documents loaded during processing")
	rootCmd.Flags().String("maxmem", "", "Test the parser memory support")
	rootCmd.Flags().Bool("memory", false, "Parse from memory")
	rootCmd.Flags().Bool("noblanks", false, "Drop ignorable blank spaces")
	rootCmd.Flags().Bool("nocatalogs", false, "Do not use any catalogs")
	rootCmd.Flags().Bool("nocdata", false, "Substitute CDATA section by equivalent text nodes")
	rootCmd.Flags().Bool("noent", false, "Substitute entity values for entity references")
	rootCmd.Flags().Bool("nonet", false, "Do not use the Internet to fetch DTDs or entities")
	rootCmd.Flags().Bool("noout", false, "Suppress output")
	rootCmd.Flags().Bool("nowarning", false, "Do not emit warnings")
	rootCmd.Flags().Bool("nowrap", false, "Do not output HTML doc wrapper")
	rootCmd.Flags().Bool("noxincludenode", false, "Do XInclude but do not generate XInclude start/end nodes")
	rootCmd.Flags().Bool("nsclean", false, "Remove redundant namespace declarations")
	rootCmd.Flags().String("output", "", "Define a file path to save the result")
	rootCmd.Flags().String("path", "", "Use the list of filesystem paths to load DTDs or entities")
	rootCmd.Flags().String("pattern", "", "Pattern recognition engine")
	rootCmd.Flags().Bool("postvalid", false, "Validate after parsing")
	rootCmd.Flags().Bool("push", false, "Use the push mode of the parser")
	rootCmd.Flags().Bool("recover", false, "Output any parsable portions of an invalid document")
	rootCmd.Flags().String("relaxng", "", "Use RelaxNG file for validation")
	rootCmd.Flags().Bool("repeat", false, "Repeat 100 times for timing")
	rootCmd.Flags().String("schema", "", "Use a W3C XML Schema file for validation")
	rootCmd.Flags().Bool("shell", false, "Run a navigating shell")
	rootCmd.Flags().Bool("stream", false, "Use streaming API")
	rootCmd.Flags().Bool("testIO", false, "Test user input/output support")
	rootCmd.Flags().Bool("timing", false, "Output timing information")
	rootCmd.Flags().Bool("version", false, "Version")
	rootCmd.Flags().Bool("xinclude", false, "Do XInclude processing")
	rootCmd.Flags().String("xpath", "", "Run an XPath expression")
	rootCmd.Flags().Bool("xmlout", false, "Output results as XML")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"encode":  carapace.ActionValues("UTF-8", "UTF-16", "ISO-8859-1", "ASCII", "US-ASCII"),
		"output":  carapace.ActionFiles(),
		"path":    carapace.ActionDirectories(),
		"pattern": carapace.ActionFiles(),
		"relaxng": carapace.ActionFiles(),
		"schema":  carapace.ActionFiles(),
	})
}
