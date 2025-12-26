package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfdetach",
	Short: "Portable Document Format (PDF) document embedded file extractor (version 3.03)",
	Long:  "https://man.archlinux.org/man/pdfdetach.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().StringS("enc", "enc", "", "output text encoding name")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().BoolS("list", "list", false, "list all embedded files")
	rootCmd.Flags().StringS("o", "o", "", "file name for the saved embedded file")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().StringS("save", "save", "", "save the specified embedded file (file number)")
	rootCmd.Flags().BoolS("saveall", "saveall", false, "save all embedded files")
	rootCmd.Flags().StringS("savefile", "savefile", "", "save the specified embedded file (file name)")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
