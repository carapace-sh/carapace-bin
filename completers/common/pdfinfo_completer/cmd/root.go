package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfinfo",
	Short: "Portable Document Format (PDF) document information extractor (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdfinfo.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().BoolS("box", "box", false, "print the page bounding boxes")
	rootCmd.Flags().BoolS("custom", "custom", false, "print both custom and standard metadata")
	rootCmd.Flags().BoolS("dests", "dests", false, "print all named destinations in the PDF")
	rootCmd.Flags().StringS("enc", "enc", "", "output text encoding name")
	rootCmd.Flags().StringS("f", "f", "", "first page to convert")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("isodates", "isodates", false, "print the dates in ISO-8601 format")
	rootCmd.Flags().BoolS("js", "js", false, "print all JavaScript in the PDF")
	rootCmd.Flags().StringS("l", "l", "", "last page to convert")
	rootCmd.Flags().BoolS("listenc", "listenc", false, "list available encodings")
	rootCmd.Flags().BoolS("meta", "meta", false, "print the document metadata (XML)")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().BoolS("rawdates", "rawdates", false, "print the undecoded date strings directly from the PDF file")
	rootCmd.Flags().BoolS("struct", "struct", false, "print the logical document structure (for tagged files)")
	rootCmd.Flags().BoolS("struct-text", "struct-text", false, "print text contents along with document structure (for tagged files)")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("url", "url", false, "print all URLs inside PDF objects (does not scan text content)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
