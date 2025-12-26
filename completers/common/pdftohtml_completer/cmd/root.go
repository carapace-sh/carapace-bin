package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdftohtml",
	Short: "program to convert PDF files into HTML, XML and PNG images",
	Long:  "https://man.archlinux.org/man/pdftohtml.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().BoolS("c", "c", false, "generate complex document")
	rootCmd.Flags().BoolS("dataurls", "dataurls", false, "use data URLs instead of external images in HTML")
	rootCmd.Flags().StringS("enc", "enc", "", "output text encoding name")
	rootCmd.Flags().StringS("f", "f", "", "first page to convert")
	rootCmd.Flags().StringS("fmt", "fmt", "", "image file format for Splash output (png or jpg)")
	rootCmd.Flags().BoolS("fontfullname", "fontfullname", false, "outputs font full name")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("hidden", "hidden", false, "output hidden text")
	rootCmd.Flags().BoolS("i", "i", false, "ignore images")
	rootCmd.Flags().StringS("l", "l", "", "last page to convert")
	rootCmd.Flags().BoolS("nodrm", "nodrm", false, "override document DRM settings")
	rootCmd.Flags().BoolS("noframes", "noframes", false, "generate no frames")
	rootCmd.Flags().BoolS("nomerge", "nomerge", false, "do not merge paragraphs")
	rootCmd.Flags().BoolS("noroundcoord", "noroundcoord", false, "do not round coordinates (with XML output only)")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("p", "p", false, "exchange .pdf links by .html")
	rootCmd.Flags().BoolS("q", "q", false, "don't print any messages or errors")
	rootCmd.Flags().BoolS("s", "s", false, "generate single document that includes all pages")
	rootCmd.Flags().BoolS("stdout", "stdout", false, "use standard output")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")
	rootCmd.Flags().StringS("wbt", "wbt", "", "word break threshold (default 10 percent)")
	rootCmd.Flags().BoolS("xml", "xml", false, "output for XML post-processing")
	rootCmd.Flags().StringS("zoom", "zoom", "", "zoom the pdf document (default 1.5)")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
