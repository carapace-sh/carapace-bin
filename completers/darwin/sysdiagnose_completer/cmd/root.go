package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysdiagnose",
	Short: "gather system-wide diagnostic information",
	Long:  "https://keith.github.io/xcode-manpages/sysdiagnose.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("A", "A", "", "Archive name")
	rootCmd.Flags().StringS("C", "C", "", "Compression type")
	rootCmd.Flags().BoolS("D", "D", false, "Detailed")
	rootCmd.Flags().BoolS("F", "F", false, "Force")
	rootCmd.Flags().BoolS("G", "G", false, "Graphics")
	rootCmd.Flags().BoolS("H", "H", false, "Print path to default output directory")
	rootCmd.Flags().BoolS("P", "P", false, "Performance")
	rootCmd.Flags().BoolS("Q", "Q", false, "Quick")
	rootCmd.Flags().BoolS("R", "R", false, "Recent only")
	rootCmd.Flags().BoolS("S", "S", false, "Sysdiagnose")
	rootCmd.Flags().StringS("V", "V", "", "Volume path")
	rootCmd.Flags().BoolS("b", "b", false, "Bundle")
	rootCmd.Flags().BoolS("d", "d", false, "Diagnostic")
	rootCmd.Flags().StringS("f", "f", "", "Results directory")
	rootCmd.Flags().BoolS("g", "g", false, "GPU")
	rootCmd.Flags().BoolS("h", "h", false, "Display help")
	rootCmd.Flags().BoolS("k", "k", false, "Keep")
	rootCmd.Flags().BoolS("n", "n", false, "No compression")
	rootCmd.Flags().BoolS("p", "p", false, "Process")
	rootCmd.Flags().BoolS("r", "r", false, "Recent")
	rootCmd.Flags().BoolS("u", "u", false, "User")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"A": carapace.ActionFiles(),
		"C": carapace.ActionValues("none", "gzip", "bzip2", "lzfse", "lzma"),
		"V": carapace.ActionFiles(),
		"f": carapace.ActionDirectories(),
	})
}
