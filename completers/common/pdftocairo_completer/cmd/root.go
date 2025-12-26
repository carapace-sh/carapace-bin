package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdftocairo",
	Short: "Portable Document Format (PDF) to PNG/JPEG/TIFF/PDF/PS/EPS/SVG using cairo",
	Long:  "https://man.archlinux.org/man/pdftocairo.1",
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
	rootCmd.Flags().StringS("antialias", "antialias", "", "set cairo antialias option")
	rootCmd.Flags().BoolS("cropbox", "cropbox", false, "use the crop box rather than media box")
	rootCmd.Flags().BoolS("duplex", "duplex", false, "enable duplex printing")
	rootCmd.Flags().BoolS("e", "e", false, "print only even pages")
	rootCmd.Flags().BoolS("eps", "eps", false, "generate Encapsulated PostScript (EPS)")
	rootCmd.Flags().BoolS("expand", "expand", false, "expand pages smaller than the paper size")
	rootCmd.Flags().StringS("f", "f", "", "first page to print")
	rootCmd.Flags().BoolS("gray", "gray", false, "generate a grayscale image file (PNG, JPEG)")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().StringS("icc", "icc", "", "ICC color profile to use")
	rootCmd.Flags().BoolS("jpeg", "jpeg", false, "generate a JPEG file")
	rootCmd.Flags().StringS("jpegopt", "jpegopt", "", "jpeg options, with format <opt1>=<val1>[,<optN>=<valN>]*")
	rootCmd.Flags().StringS("l", "l", "", "last page to print")
	rootCmd.Flags().BoolS("level2", "level2", false, "generate Level 2 PostScript (PS, EPS)")
	rootCmd.Flags().BoolS("level3", "level3", false, "generate Level 3 PostScript (PS, EPS)")
	rootCmd.Flags().BoolS("mono", "mono", false, "generate a monochrome image file (PNG, JPEG)")
	rootCmd.Flags().BoolS("nocenter", "nocenter", false, "don't center pages smaller than the paper size")
	rootCmd.Flags().BoolS("nocrop", "nocrop", false, "don't crop pages to CropBox")
	rootCmd.Flags().BoolS("noshrink", "noshrink", false, "don't shrink pages larger than the paper size")
	rootCmd.Flags().BoolS("o", "o", false, "print only odd pages")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("origpagesizes", "origpagesizes", false, "conserve original page sizes (PS, PDF, SVG)")
	rootCmd.Flags().StringS("paper", "paper", "", "paper size (letter, legal, A4, A3, match)")
	rootCmd.Flags().StringS("paperh", "paperh", "", "paper height, in points")
	rootCmd.Flags().StringS("paperw", "paperw", "", "paper width, in points")
	rootCmd.Flags().BoolS("pdf", "pdf", false, "generate a PDF file")
	rootCmd.Flags().BoolS("png", "png", false, "generate a PNG file")
	rootCmd.Flags().BoolS("print", "print", false, "print to a Windows printer")
	rootCmd.Flags().BoolS("printdlg", "printdlg", false, "show Windows print dialog and print")
	rootCmd.Flags().StringS("printer", "printer", "", "printer name or use default if this option is not specified")
	rootCmd.Flags().StringS("printopt", "printopt", "", "printer options, with format <opt1>=<val1>[,<optN>=<valN>]*")
	rootCmd.Flags().BoolS("ps", "ps", false, "generate PostScript file")
	rootCmd.Flags().BoolS("q", "q", false, "don't print any messages or errors")
	rootCmd.Flags().StringS("r", "r", "", "resolution, in PPI (default is 150)")
	rootCmd.Flags().StringS("rx", "rx", "", "X resolution, in PPI (default is 150)")
	rootCmd.Flags().StringS("ry", "ry", "", "Y resolution, in PPI (default is 150)")
	rootCmd.Flags().StringS("scale-to", "scale-to", "", "scales each page to fit within scale-to*scale-to pixel box")
	rootCmd.Flags().StringS("scale-to-x", "scale-to-x", "", "scales each page horizontally to fit in scale-to-x pixels")
	rootCmd.Flags().StringS("scale-to-y", "scale-to-y", "", "scales each page vertically to fit in scale-to-y pixels")
	rootCmd.Flags().BoolS("setupdlg", "setupdlg", false, "show printer setup dialog before printing")
	rootCmd.Flags().BoolS("singlefile", "singlefile", false, "write only the first page and do not add digits")
	rootCmd.Flags().BoolS("struct", "struct", false, "enable logical document structure")
	rootCmd.Flags().BoolS("svg", "svg", false, "generate a Scalable Vector Graphics (SVG) file")
	rootCmd.Flags().StringS("sz", "sz", "", "size of crop square in pixels (sets W and H)")
	rootCmd.Flags().BoolS("tiff", "tiff", false, "generate a TIFF file")
	rootCmd.Flags().StringS("tiffcompression", "tiffcompression", "", "set TIFF compression")
	rootCmd.Flags().BoolS("transp", "transp", false, "use a transparent background instead of white (PNG)")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")
	rootCmd.Flags().StringS("x", "x", "", "x-coordinate of the crop area top left corner")
	rootCmd.Flags().StringS("y", "y", "", "y-coordinate of the crop area top left corner")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"tiffcompression": carapace.ActionValues("none", "packbits", "jpeg", "lzw", "deflate"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
