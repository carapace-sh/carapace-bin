package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdftoppm",
	Short: "Portable Document Format (PDF) to Portable Pixmap (PPM) converter (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdftoppm.1",
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
	rootCmd.Flags().StringS("aa", "aa", "", "enable font anti-aliasing")
	rootCmd.Flags().StringS("aaVector", "aaVector", "", "enable vector anti-aliasing")
	rootCmd.Flags().BoolS("cropbox", "cropbox", false, "use the crop box rather than media box")
	rootCmd.Flags().StringS("defaultcmykprofile", "defaultcmykprofile", "", "ICC color profile to use as the DefaultCMYK color space")
	rootCmd.Flags().StringS("defaultgrayprofile", "defaultgrayprofile", "", "ICC color profile to use as the DefaultGray color space")
	rootCmd.Flags().StringS("defaultrgbprofile", "defaultrgbprofile", "", "ICC color profile to use as the DefaultRGB color space")
	rootCmd.Flags().StringS("displayprofile", "displayprofile", "", "ICC color profile to use as the display profile")
	rootCmd.Flags().BoolS("e", "e", false, "print only even pages")
	rootCmd.Flags().StringS("f", "f", "", "first page to print")
	rootCmd.Flags().BoolS("forcenum", "forcenum", false, "force page number even if there is only one page")
	rootCmd.Flags().StringS("freetype", "freetype", "", "enable FreeType font rasterizer")
	rootCmd.Flags().BoolS("gray", "gray", false, "generate a grayscale PGM file")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("hide-annotations", "hide-annotations", false, "do not show annotations")
	rootCmd.Flags().BoolS("jpeg", "jpeg", false, "generate a JPEG file")
	rootCmd.Flags().BoolS("jpegcmyk", "jpegcmyk", false, "generate a CMYK JPEG file")
	rootCmd.Flags().StringS("jpegopt", "jpegopt", "", "jpeg options, with format <opt1>=<val1>[,<optN>=<valN>]*")
	rootCmd.Flags().StringS("l", "l", "", "last page to print")
	rootCmd.Flags().BoolS("mono", "mono", false, "generate a monochrome PBM file")
	rootCmd.Flags().BoolS("o", "o", false, "print only odd pages")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("overprint", "overprint", false, "enable overprint")
	rootCmd.Flags().BoolS("png", "png", false, "generate a PNG file")
	rootCmd.Flags().BoolS("progress", "progress", false, "print progress info")
	rootCmd.Flags().BoolS("q", "q", false, "don't print any messages or errors")
	rootCmd.Flags().StringS("r", "r", "", "resolution, in DPI (default is 150)")
	rootCmd.Flags().StringS("rx", "rx", "", "X resolution, in DPI (default is 150)")
	rootCmd.Flags().StringS("ry", "ry", "", "Y resolution, in DPI (default is 150)")
	rootCmd.Flags().BoolS("scale-dimension-before-rotation", "scale-dimension-before-rotation", false, "for rotated pdf, resize dimensions before the rotation")
	rootCmd.Flags().StringS("scale-to", "scale-to", "", "scales each page to fit within scale-to*scale-to pixel box")
	rootCmd.Flags().StringS("scale-to-x", "scale-to-x", "", "scales each page horizontally to fit in scale-to-x pixels")
	rootCmd.Flags().StringS("scale-to-y", "scale-to-y", "", "scales each page vertically to fit in scale-to-y pixels")
	rootCmd.Flags().StringS("sep", "sep", "", "single character separator between name and page number, default -")
	rootCmd.Flags().BoolS("singlefile", "singlefile", false, "write only the first page and do not add digits")
	rootCmd.Flags().StringS("sz", "sz", "", "size of crop square in pixels (sets W and H)")
	rootCmd.Flags().StringS("thinlinemode", "thinlinemode", "", "set thin line mode")
	rootCmd.Flags().BoolS("tiff", "tiff", false, "generate a TIFF file")
	rootCmd.Flags().StringS("tiffcompression", "tiffcompression", "", "set TIFF compression")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")
	rootCmd.Flags().StringS("x", "x", "", "x-coordinate of the crop area top left corner")
	rootCmd.Flags().StringS("y", "y", "", "y-coordinate of the crop area top left corner")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"aa":              carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"aaVector":        carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"freetype":        carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"thinlinemode":    carapace.ActionValues("none", "solid", "shape"),
		"tiffcompression": carapace.ActionValues("none", "packbits", "jpeg", "lzw", "deflate"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
