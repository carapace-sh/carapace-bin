package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jot",
	Short: "print sequential or random data",
	Long:  "https://man.freebsd.org/cgi/man.cgi?jot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("b", "b", "", "Word to print repetitively")
	rootCmd.Flags().BoolS("c", "c", false, "Character format")
	rootCmd.Flags().BoolS("n", "n", false, "Do not print final newline")
	rootCmd.Flags().StringS("p", "p", "", "Precision")
	rootCmd.Flags().BoolS("r", "r", false, "Generate random data")
	rootCmd.Flags().StringS("s", "s", "", "String separator")
	rootCmd.Flags().StringS("w", "w", "", "Word with data appended")
}
