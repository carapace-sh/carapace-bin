package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdftops",
	Short: "Portable Document Format (PDF) to PostScript converter (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdftops.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().StringS("aaRaster", "aaRaster", "", "enable anti-aliasing on rasterization")
	rootCmd.Flags().BoolS("binary", "binary", false, "write binary data in Level 1 PostScript")
	rootCmd.Flags().StringS("defaultcmykprofile", "defaultcmykprofile", "", "ICC color profile to use as the DefaultCMYK color space")
	rootCmd.Flags().StringS("defaultgrayprofile", "defaultgrayprofile", "", "ICC color profile to use as the DefaultGray color space")
	rootCmd.Flags().StringS("defaultrgbprofile", "defaultrgbprofile", "", "ICC color profile to use as the DefaultRGB color space")
	rootCmd.Flags().BoolS("duplex", "duplex", false, "enable duplex printing")
	rootCmd.Flags().BoolS("eps", "eps", false, "generate Encapsulated PostScript (EPS)")
	rootCmd.Flags().BoolS("expand", "expand", false, "expand pages smaller than the paper size")
	rootCmd.Flags().StringS("f", "f", "", "first page to print")
	rootCmd.Flags().BoolS("form", "form", false, "generate a PostScript form")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().StringS("l", "l", "", "last page to print")
	rootCmd.Flags().BoolS("level1", "level1", false, "generate Level 1 PostScript")
	rootCmd.Flags().BoolS("level1sep", "level1sep", false, "generate Level 1 separable PostScript")
	rootCmd.Flags().BoolS("level2", "level2", false, "generate Level 2 PostScript")
	rootCmd.Flags().BoolS("level2sep", "level2sep", false, "generate Level 2 separable PostScript")
	rootCmd.Flags().BoolS("level3", "level3", false, "generate Level 3 PostScript")
	rootCmd.Flags().BoolS("level3sep", "level3sep", false, "generate Level 3 separable PostScript")
	rootCmd.Flags().BoolS("nocenter", "nocenter", false, "don't center pages smaller than the paper size")
	rootCmd.Flags().BoolS("nocrop", "nocrop", false, "don't crop pages to CropBox")
	rootCmd.Flags().BoolS("noembcidps", "noembcidps", false, "don't embed CID PostScript fonts")
	rootCmd.Flags().BoolS("noembcidtt", "noembcidtt", false, "don't embed CID TrueType fonts")
	rootCmd.Flags().BoolS("noembt1", "noembt1", false, "don't embed Type 1 fonts")
	rootCmd.Flags().BoolS("noembtt", "noembtt", false, "don't embed TrueType fonts")
	rootCmd.Flags().BoolS("noshrink", "noshrink", false, "don't shrink pages larger than the paper size")
	rootCmd.Flags().BoolS("opi", "opi", false, "generate OPI comments")
	rootCmd.Flags().BoolS("optimizecolorspace", "optimizecolorspace", false, "convert gray RGB images to gray color space")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("origpagesizes", "origpagesizes", false, "conserve original page sizes")
	rootCmd.Flags().BoolS("overprint", "overprint", false, "enable overprint emulation during rasterization")
	rootCmd.Flags().StringS("paper", "paper", "", "paper size (letter, legal, A4, A3, match)")
	rootCmd.Flags().StringS("paperh", "paperh", "", "paper height, in points")
	rootCmd.Flags().StringS("paperw", "paperw", "", "paper width, in points")
	rootCmd.Flags().BoolS("passfonts", "passfonts", false, "don't substitute missing fonts")
	rootCmd.Flags().BoolS("passlevel1customcolor", "passlevel1customcolor", false, "pass custom color in level1sep")
	rootCmd.Flags().BoolS("preload", "preload", false, "preload images and forms")
	rootCmd.Flags().StringS("processcolorformat", "processcolorformat", "", "color format that is used during rasterization and transparency reduction")
	rootCmd.Flags().StringS("processcolorprofile", "processcolorprofile", "", "ICC color profile to use as the process color profile during rasterization and transparency reduction")
	rootCmd.Flags().BoolS("q", "q", false, "don't print any messages or errors")
	rootCmd.Flags().StringS("r", "r", "", "resolution for rasterization, in DPI (default is 300)")
	rootCmd.Flags().StringS("rasterize", "rasterize", "", "control rasterization")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"aaRaster":           carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"processcolorformat": carapace.ActionValues("MONO8", "RGB8", "CMYK8"),
		"rasterize":          carapace.ActionValues("always", "never", "whenneeded"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
