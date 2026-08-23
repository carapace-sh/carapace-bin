package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atos",
	Short: "convert numeric addresses to symbols",
	Long:  "https://keith.github.io/xcode-manpages/atos.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "Print the disassembly")
	rootCmd.Flags().BoolS("f", "f", false, "Print the full path")
	rootCmd.Flags().StringS("l", "l", "", "Specify the address of the library start")
	rootCmd.Flags().StringS("o", "o", "", "Load addresses from specified file")
	rootCmd.Flags().BoolS("p", "p", false, "Print the process name")
	rootCmd.Flags().BoolS("s", "s", false, "Print the shared library name")
	rootCmd.Flags().BoolS("v", "v", false, "Print the version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionFiles(),
	})
}
