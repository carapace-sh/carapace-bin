package cmd

import (
	"github.com/carapace-sh/carapace-bin/completers/gdown_completer/cmd/actions"
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gdown",
	Short: "Google Drive Public File Downloader when Curl/Wget Fails",
	Long:  "https://github.com/wkentaro/gdown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("continue", "c", false, "resume getting partially-downloaded files")
	rootCmd.Flags().Bool("folder", false, "download entire folder instead of a single file")
	rootCmd.Flags().String("format", "", "Format of Google Docs, Spreadsheets, and Slides")
	rootCmd.Flags().Bool("fuzzy", false, "extract Google Drive's file ID")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("no-check-certificate", false, "don't check the server's TLS certificate")
	rootCmd.Flags().Bool("no-cookies", false, "don't use cookies in ~/.cache/gdown/cookies.txt")
	rootCmd.Flags().StringP("output", "O", "", "output file name/path")
	rootCmd.Flags().String("proxy", "", "download using the specified proxy")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress logging except errors")
	rootCmd.Flags().Bool("remaining-ok", false, "asserts that is ok to download max 50 files per folder")
	rootCmd.Flags().String("speed", "", "download speed limit in second")
	rootCmd.Flags().String("user-agent", "", "User-Agent to use for downloading file")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"docx", "Google Docs",
			"xlsx", "Spreadsheet",
			"pptx", "Slides",
		),
		"speed":  actions.ActionSpeed(),
		"output": carapace.ActionFiles(),
	})
}
