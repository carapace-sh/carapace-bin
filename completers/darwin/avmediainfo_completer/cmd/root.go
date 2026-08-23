package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "avmediainfo",
	Short: "display media file metadata",
	Long:  "https://keith.github.io/xcode-manpages/avmediainfo.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("format", false, "Display format")
	rootCmd.Flags().String("formatType", "", "Format type")
	rootCmd.Flags().Bool("help", false, "Show help")
	rootCmd.Flags().Bool("metadata", false, "Display metadata")
	rootCmd.Flags().Bool("tracks", false, "Display tracks")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
