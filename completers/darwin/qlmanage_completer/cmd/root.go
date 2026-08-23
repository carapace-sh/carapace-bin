package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qlmanage",
	Short: "Quick Look Server debug and management tool",
	Long:  "https://keith.github.io/xcode-manpages/qlmanage.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("c", "", "Specify content type UTI")
	rootCmd.Flags().String("f", "", "Specify scale factor")
	rootCmd.Flags().String("g", "", "Specify generator")
	rootCmd.Flags().BoolP("help", "h", false, "Display extensive help")
	rootCmd.Flags().Bool("i", false, "Use icon mode")
	rootCmd.Flags().BoolP("info", "m", false, "Get information on Quick Look server")
	rootCmd.Flags().BoolP("preview", "p", false, "Display generated previews for the specified files")
	rootCmd.Flags().BoolP("reset", "r", false, "Reset Quick Look Server and all generator caches")
	rootCmd.Flags().String("s", "", "Specify thumbnail size")
	rootCmd.Flags().BoolP("thumbnail", "t", false, "Display generated thumbnails for the specified files")
	rootCmd.Flags().Bool("x", false, "Use XML output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionValues(),
		"g": carapace.ActionValues(),
		"s": carapace.ActionValues("16", "32", "64", "128", "256", "512", "1024"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}