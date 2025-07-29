package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dngconverter",
	Short: "Adobe DNG Converter",
	Long:  "https://helpx.adobe.com/camera-raw/using/adobe-dng-converter.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "output lossless compressed DNG files")
	rootCmd.Flags().StringS("count", "count", "", "limit pixel count to <num> pixels/image")
	rootCmd.Flags().BoolS("cr11.2", "cr11.2", false, "set Camera Raw compatibility to 11.2 and later")
	rootCmd.Flags().BoolS("cr12.4", "cr12.4", false, "set Camera Raw compatibility to 12.4 and later")
	rootCmd.Flags().BoolS("cr13.2", "cr13.2", false, "set Camera Raw compatibility to 13.2 and later")
	rootCmd.Flags().BoolS("cr14.0", "cr14.0", false, "set Camera Raw compatibility to 14.0 and later")
	rootCmd.Flags().BoolS("cr15.3", "cr15.3", false, "set Camera Raw compatibility to 15.3 and later")
	rootCmd.Flags().BoolS("cr2.4", "cr2.4", false, "set Camera Raw compatibility to 2.4 and later")
	rootCmd.Flags().BoolS("cr4.1", "cr4.1", false, "set Camera Raw compatibility to 4.1 and later")
	rootCmd.Flags().BoolS("cr4.6", "cr4.6", false, "set Camera Raw compatibility to 4.6 and later")
	rootCmd.Flags().BoolS("cr5.4", "cr5.4", false, "set Camera Raw compatibility to 5.4 and later")
	rootCmd.Flags().BoolS("cr6.6", "cr6.6", false, "set Camera Raw compatibility to 6.6 and later")
	rootCmd.Flags().BoolS("cr7.1", "cr7.1", false, "set Camera Raw compatibility to 7.1 and later")
	rootCmd.Flags().StringS("d", "d", "", "output converted files to the specified directory")
	rootCmd.Flags().BoolS("dng1.1", "dng1.1", false, "set DNG backward version to 1.1")
	rootCmd.Flags().BoolS("dng1.3", "dng1.3", false, "set DNG backward version to 1.3")
	rootCmd.Flags().BoolS("dng1.4", "dng1.4", false, "set DNG backward version to 1.4")
	rootCmd.Flags().BoolS("dng1.5", "dng1.5", false, "set DNG backward version to 1.5")
	rootCmd.Flags().BoolS("dng1.6", "dng1.6", false, "set DNG backward version to 1.6")
	rootCmd.Flags().BoolS("dng1.7", "dng1.7", false, "set DNG backward version to 1.7")
	rootCmd.Flags().BoolS("dng1.7.1", "dng1.7.1", false, "set DNG backward version to 1.7.1")
	rootCmd.Flags().BoolS("e", "e", false, "embed original raw file inside DNG files")
	rootCmd.Flags().BoolS("fl", "fl", false, "embed fast load data inside DNG files")
	rootCmd.Flags().BoolS("jxl", "jxl", false, "use JPEG XL compression, if supported by image type")
	rootCmd.Flags().StringS("jxl_distance", "jxl_distance", "", "set JPEG XL distance metric")
	rootCmd.Flags().StringS("jxl_effort", "jxl_effort", "", "set JPEG XL effort level")
	rootCmd.Flags().BoolS("l", "l", false, "output linear DNG files")
	rootCmd.Flags().BoolS("losslessJXL", "losslessJXL", false, "uses Lossless JPEG XL compression")
	rootCmd.Flags().BoolS("lossy", "lossy", false, "use lossy compression")
	rootCmd.Flags().BoolS("lossyMosaicJXL", "lossyMosaicJXL", false, "uses Lossy JPEG XL compression with Bayer images")
	rootCmd.Flags().BoolS("mp", "mp", false, "process multiple raw files in parallel")
	rootCmd.Flags().StringS("o", "o", "", "specify the name of the output DNG file")
	rootCmd.Flags().BoolS("p0", "p0", false, "set JPEG preview size to none")
	rootCmd.Flags().BoolS("p1", "p1", false, "set JPEG preview size to medium size (default)")
	rootCmd.Flags().BoolS("p2", "p2", false, "set JPEG preview size to full size")
	rootCmd.Flags().StringS("side", "side", "", "limit size to <num> pixels/side")
	rootCmd.Flags().BoolS("u", "u", false, "output uncompressed DNG files")

	rootCmd.MarkFlagsMutuallyExclusive(
		"c",
		"u",
		"l",
	)

	rootCmd.MarkFlagsMutuallyExclusive(
		"cr2.4",
		"cr4.1",
		"cr4.6",
		"cr5.4",
		"cr6.6",
		"cr7.1",
		"cr11.2",
		"cr12.4",
		"cr13.2",
		"cr14.0",
		"cr15.3",
	)

	rootCmd.MarkFlagsMutuallyExclusive(
		"dng1.1",
		"dng1.3",
		"dng1.4",
		"dng1.5",
		"dng1.6",
		"dng1.7",
		"dng1.7.1",
	)

	rootCmd.MarkFlagsMutuallyExclusive(
		"p0",
		"p1",
		"p2",
	)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"d": carapace.ActionDirectories(),
		"jxl_distance": carapace.ActionValuesDescribed(
			"0.0", "lossless",
			"1.0", "",
			"2.0", "",
			"3.0", "",
			"4.0", "",
			"5.0", "",
			"6.0", "lossy",
		),
		"jxl_effort": carapace.ActionValuesDescribed(
			"1", "fastest",
			"2", "",
			"3", "",
			"4", "",
			"5", "",
			"6", "",
			"7", "",
			"8", "",
			"9", "slowest",
		),
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
