package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdftotext",
	Short: "Portable Document Format (PDF) to text converter (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdftotext.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().StringS("H", "H", "", "height of crop area in pixels (default is 0)")
	rootCmd.Flags().StringS("W", "W", "", "width of crop area in pixels (default is 0)")
	rootCmd.Flags().BoolS("bbox", "bbox", false, "output bounding box for each word and page size to html. Sets -htmlmeta")
	rootCmd.Flags().BoolS("bbox-layout", "bbox-layout", false, "like -bbox but with extra layout bounding box data.  Sets -htmlmeta")
	rootCmd.Flags().StringS("colspacing", "colspacing", "", "how much spacing we allow after a word before considering adjacent text to be a new column, as a fraction of the font size (default is 0.7, old releases had a 0.3 default)")
	rootCmd.Flags().BoolS("cropbox", "cropbox", false, "use the crop box rather than media box")
	rootCmd.Flags().StringS("enc", "enc", "", "output text encoding name")
	rootCmd.Flags().StringS("eol", "eol", "", "output end-of-line convention (unix, dos, or mac)")
	rootCmd.Flags().StringS("f", "f", "", "first page to convert")
	rootCmd.Flags().StringS("fixed", "fixed", "", "assume fixed-pitch (or tabular) text")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("htmlmeta", "htmlmeta", false, "generate a simple HTML file, including the meta information")
	rootCmd.Flags().StringS("l", "l", "", "last page to convert")
	rootCmd.Flags().BoolS("layout", "layout", false, "maintain original physical layout")
	rootCmd.Flags().BoolS("listenc", "listenc", false, "list available encodings")
	rootCmd.Flags().BoolS("nodiag", "nodiag", false, "discard diagonal text")
	rootCmd.Flags().BoolS("nopgbrk", "nopgbrk", false, "don't insert page breaks between pages")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("q", "q", false, "don't print any messages or errors")
	rootCmd.Flags().StringS("r", "r", "", "resolution, in DPI (default is 72)")
	rootCmd.Flags().BoolS("raw", "raw", false, "keep strings in content stream order")
	rootCmd.Flags().BoolS("tsv", "tsv", false, "generate a simple TSV file, including the meta information for bounding boxes")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")
	rootCmd.Flags().StringS("x", "x", "", "x-coordinate of the crop area top left corner")
	rootCmd.Flags().StringS("y", "y", "", "y-coordinate of the crop area top left corner")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"eol": carapace.ActionValues("unix", "dos", "mac"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
