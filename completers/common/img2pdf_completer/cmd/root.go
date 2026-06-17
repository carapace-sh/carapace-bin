package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/img2pdf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "img2pdf",
	Short: "losslessly convert raster images to PDF",
	Long:  "https://gitlab.mister-muffin.de/josch/img2pdf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("art-border", "", "border between the ArtBox and the MediaBox")
	rootCmd.Flags().String("author", "", "sets the author metadata value")
	rootCmd.Flags().BoolP("auto-orient", "a", false, "conditionally swap page dimensions to match input image orientation")
	rootCmd.Flags().String("bleed-border", "", "border between the BleedBox and the MediaBox")
	rootCmd.Flags().StringP("border", "b", "", "minimal distance between image border and PDF page border")
	rootCmd.Flags().StringP("colorspace", "C", "", "forces the PIL colorspace")
	rootCmd.Flags().String("creationdate", "", "sets the UTC creation date metadata value")
	rootCmd.Flags().String("creator", "", "sets the creator metadata value")
	rootCmd.Flags().String("crop-border", "", "border between the CropBox and the MediaBox")
	rootCmd.Flags().String("engine", "", "choose PDF engine (internal, pikepdf or pdfrw)")
	rootCmd.Flags().Bool("first-frame-only", false, "only convert the first frame of multi-frame input images")
	rootCmd.Flags().StringP("fit", "f", "into", "fit the image using the given dimensions")
	rootCmd.Flags().String("from-file", "", "read the list of images from FILE")
	rootCmd.Flags().Bool("gui", false, "run experimental tkinter gui")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().StringP("imgsize", "s", "", "sets the size of the images on the PDF pages")
	rootCmd.Flags().Bool("include-thumbnails", false, "create one page per frame including thumbnails")
	rootCmd.Flags().StringArray("keywords", nil, "sets the keywords metadata value")
	rootCmd.Flags().String("moddate", "", "sets the UTC modification date metadata value")
	rootCmd.Flags().BoolP("nodate", "D", false, "suppress timestamps in the output")
	rootCmd.Flags().String("orientation", "auto", "alias for --rotation")
	rootCmd.Flags().StringP("output", "o", "", "output to a file instead of standard output")
	rootCmd.Flags().StringP("pagesize", "S", "", "sets the size of the PDF pages")
	rootCmd.Flags().String("pdfa", "", "output a PDF/A-1b compliant document")
	rootCmd.Flags().Bool("pillow-limit-break", false, "disable Pillow decompression bomb size limit")
	rootCmd.Flags().String("producer", "", "sets the producer metadata value")
	rootCmd.Flags().StringP("rotation", "r", "auto", "specifies how input images should be rotated")
	rootCmd.Flags().String("subject", "", "sets the subject metadata value")
	rootCmd.Flags().String("title", "", "sets the title metadata value")
	rootCmd.Flags().String("trim-border", "", "border between the TrimBox and the MediaBox")
	rootCmd.Flags().BoolP("verbose", "v", false, "operate in verbose mode")
	rootCmd.Flags().BoolP("version", "V", false, "prints version information and exits")
	rootCmd.Flags().Bool("viewer-center-window", false, "instruct the PDF viewer to center the window")
	rootCmd.Flags().Bool("viewer-fit-window", false, "instruct the PDF viewer to resize the window to fit the page size")
	rootCmd.Flags().Bool("viewer-fullscreen", false, "instruct the PDF viewer to open in fullscreen mode")
	rootCmd.Flags().Int("viewer-initial-page", 0, "instruct the PDF viewer to show the given page")
	rootCmd.Flags().String("viewer-magnification", "", "instruct the PDF viewer to open with a certain zoom level")
	rootCmd.Flags().String("viewer-page-layout", "", "instruct the PDF viewer how to arrange the pages on the screen")
	rootCmd.Flags().String("viewer-panes", "", "instruct the PDF viewer which side panes to show")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"author": carapace.ActionValues(),
		"colorspace": carapace.ActionValuesDescribed(
			"RGB", "RGB color",
			"L", "Grayscale",
			"1", "Black and white (internally converted to grayscale)",
			"CMYK", "CMYK color",
			"CMYK;I", "CMYK color with inversion (for CMYK JPEG files from Adobe)",
		),
		"creationdate": carapace.ActionValues(),
		"creator":      carapace.ActionValues(),
		"engine":       carapace.ActionValues("internal", "pikepdf", "pdfrw"),
		"fit": carapace.ActionValuesDescribed(
			"into", "width and height specify maximum values",
			"fill", "width and height specify minimum values",
			"exact", "width and height emphatically given",
			"shrink", "shrinks larger images (otherwise behaves like into)",
			"enlarge", "enlarges smaller images (otherwise behaves like into)",
		),
		"from-file":            carapace.ActionFiles(),
		"imgsize":              action.ActionPaperSizes(),
		"keywords":             carapace.ActionValues(),
		"moddate":              carapace.ActionValues(),
		"orientation":          action.ActionRotations(),
		"output":               carapace.ActionFiles(),
		"pagesize":             action.ActionPaperSizes(),
		"pdfa":                 carapace.ActionFiles(),
		"producer":             carapace.ActionValues(),
		"rotation":             action.ActionRotations(),
		"subject":              carapace.ActionValues(),
		"title":                carapace.ActionValues(),
		"viewer-initial-page":  carapace.ActionValues(),
		"viewer-magnification": carapace.ActionValues("fit", "fith", "fitbh"),
		"viewer-page-layout":   carapace.ActionValues("single", "onecolumn", "twocolumnright", "twocolumnleft", "twopageright", "twopageleft"),
		"viewer-panes":         carapace.ActionValues("outlines", "thumbs"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
