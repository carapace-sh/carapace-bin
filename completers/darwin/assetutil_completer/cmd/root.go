package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "assetutil",
	Short: "process asset catalog .car files",
	Long:  "https://man.freebsd.org/cgi/man.cgi?assetutil",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("I", "I", false, "JSON description of asset catalog object")
	rootCmd.Flags().BoolS("M", "M", false, "Keep assets with memory class")
	rootCmd.Flags().BoolS("T", "T", false, "Compare thinning attributes")
	rootCmd.Flags().BoolS("V", "V", false, "Version")
	rootCmd.Flags().BoolS("Z", "Z", false, "Integrity check")
	rootCmd.Flags().StringS("c", "c", "", "Main Assets.car file")
	rootCmd.Flags().StringS("g", "g", "", "Graphics class")
	rootCmd.Flags().BoolS("h", "h", false, "Process hosted idioms list")
	rootCmd.Flags().StringS("i", "i", "", "Idiom to keep")
	rootCmd.Flags().StringS("n", "n", "", "Names to remove")
	rootCmd.Flags().StringS("o", "o", "", "Output file name")
	rootCmd.Flags().StringS("p", "p", "", "Display gamut")
	rootCmd.Flags().StringS("s", "s", "", "Scale factor")
	rootCmd.Flags().StringS("t", "t", "", "Subtype")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"i": carapace.ActionValues("universal", "phone", "pad", "tv", "car", "watch", "mac"),
		"o": carapace.ActionFiles(),
	})
}
