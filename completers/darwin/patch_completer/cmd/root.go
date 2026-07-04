package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "patch",
	Short: "apply a diff file to an original",
	Long:  "https://keith.github.io/xcode-manpages/patch.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Check patch")
	rootCmd.Flags().BoolS("E", "E", false, "Remove empty files")
	rootCmd.Flags().BoolS("N", "N", false, "Forward-only")
	rootCmd.Flags().BoolS("R", "R", false, "Reverse patch")
	rootCmd.Flags().BoolS("b", "b", false, "Save backups")
	rootCmd.Flags().BoolS("c", "c", false, "Context diff")
	rootCmd.Flags().BoolS("e", "e", false, "Ed script")
	rootCmd.Flags().BoolS("f", "f", false, "Force")
	rootCmd.Flags().BoolS("l", "l", false, "Ignore whitespace")
	rootCmd.Flags().BoolS("n", "n", false, "Normal diff")
	rootCmd.Flags().BoolS("s", "s", false, "Quiet")
	rootCmd.Flags().BoolS("t", "t", false, "Batch mode")
	rootCmd.Flags().BoolS("u", "u", false, "Unified diff")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")

	rootCmd.Flags().StringS("B", "B", "", "Backup prefix")
	rootCmd.Flags().StringS("D", "D", "", "Define symbol")
	rootCmd.Flags().StringS("F", "F", "", "Max fuzz factor")
	rootCmd.Flags().StringS("V", "V", "", "Version control")
	rootCmd.Flags().StringS("d", "d", "", "Change to directory")
	rootCmd.Flags().StringS("i", "i", "", "Input patchfile")
	rootCmd.Flags().StringS("o", "o", "", "Output file")
	rootCmd.Flags().StringS("p", "p", "", "Strip count")
	rootCmd.Flags().StringS("r", "r", "", "Reject file name")
	rootCmd.Flags().StringS("x", "x", "", "Debug level")
	rootCmd.Flags().StringS("z", "z", "", "Backup extension")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"i": carapace.ActionFiles(),
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
