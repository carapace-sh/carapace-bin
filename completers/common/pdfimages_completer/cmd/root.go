package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfimages",
	Short: "Portable Document Format (PDF) image extractor (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdfimages.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().BoolS("all", "all", false, "equivalent to -png -tiff -j -jp2 -jbig2 -ccitt")
	rootCmd.Flags().BoolS("ccitt", "ccitt", false, "write CCITT images as CCITT files")
	rootCmd.Flags().StringS("f", "f", "", "first page to convert")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("j", "j", false, "write JPEG images as JPEG files")
	rootCmd.Flags().BoolS("jbig2", "jbig2", false, "write JBIG2 images as JBIG2 files")
	rootCmd.Flags().BoolS("jp2", "jp2", false, "write JPEG2000 images as JP2 files")
	rootCmd.Flags().StringS("l", "l", "", "last page to convert")
	rootCmd.Flags().BoolS("list", "list", false, "print list of images instead of saving")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("p", "p", false, "include page numbers in output file names")
	rootCmd.Flags().BoolS("png", "png", false, "change the default output format to PNG")
	rootCmd.Flags().BoolS("print-filenames", "print-filenames", false, "print image filenames to stdout")
	rootCmd.Flags().BoolS("q", "q", false, "don't print any messages or errors")
	rootCmd.Flags().BoolS("tiff", "tiff", false, "change the default output format to TIFF")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
