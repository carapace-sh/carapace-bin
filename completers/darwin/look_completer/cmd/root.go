package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "look",
	Short: "display lines beginning with a given string",
	Long:  "https://man.freebsd.org/cgi/man.cgi?look",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "Dictionary character set and order")
	rootCmd.Flags().BoolS("f", "f", false, "Ignore case")
	rootCmd.Flags().StringS("t", "t", "", "Specify a string termination character")
}
