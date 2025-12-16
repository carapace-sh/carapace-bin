package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pngcheck",
	Short: "Test PNG, JNG or MNG image files for corruption",
	Long:  "http://www.libpng.org/pub/png/apps/pngcheck.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolS("7", "7", false, "print contents of tEXt chunks, escape chars >=128 (for 7-bit terminals)")
	rootCmd.Flags().BoolS("c", "c", false, "colorize output (for ANSI terminals)")
	rootCmd.Flags().BoolS("p", "p", false, "print contents of PLTE, tRNS, hIST, sPLT and PPLT (can be used with -q)")
	rootCmd.Flags().BoolS("q", "q", false, "test quietly (output only errors)")
	rootCmd.Flags().BoolS("s", "s", false, "search for PNGs within another file")
	rootCmd.Flags().BoolS("t", "t", false, "print contents of tEXt chunks (can be used with -q)")
	rootCmd.Flags().CountS("v", "v", "test verbosely (print most chunk data)")
	rootCmd.Flags().BoolS("w", "w", false, "suppress windowBits test (a more-stringent compression check)")
	rootCmd.Flags().BoolS("x", "x", false, "search for PNGs within another file and extract them when found")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".png", ".jng", ".mng"),
	)
}
