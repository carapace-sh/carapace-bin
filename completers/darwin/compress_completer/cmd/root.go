package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress data",
	Long:  "https://keith.github.io/xcode-manpages/compress.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("b", "b", "", "Limit code size to bits")
	rootCmd.Flags().BoolS("c", "c", false, "Write to stdout")
	rootCmd.Flags().BoolS("f", "f", false, "Force overwrite without prompting")
	rootCmd.Flags().BoolS("v", "v", false, "Print the percentage reduction of each file")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}