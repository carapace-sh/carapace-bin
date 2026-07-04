package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "base64",
	Short: "encode and decode using Base64 representation",
	Long:  "https://keith.github.io/xcode-manpages/base64.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Decode input (same as -d)")
	rootCmd.Flags().StringS("b", "b", "", "Break encoded lines at count characters")
	rootCmd.Flags().BoolS("d", "d", false, "Decode input")
	rootCmd.Flags().BoolS("h", "h", false, "Display help")
	rootCmd.Flags().StringS("i", "i", "", "Input file")
	rootCmd.Flags().BoolS("m", "m", false, "Use Base64 encoding method (not traditional uuencode)")
	rootCmd.Flags().StringS("o", "o", "", "Output file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"i": carapace.ActionFiles(),
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
