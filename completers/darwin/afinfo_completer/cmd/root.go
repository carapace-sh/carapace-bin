package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "afinfo",
	Short: "Audio File Info",
	Long:  "https://man.freebsd.org/cgi/man.cgi?afinfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "Brief info")
	rootCmd.Flags().StringS("c", "c", "", "Find and print a chunk")
	rootCmd.Flags().BoolS("h", "h", false, "Print help text")
	rootCmd.Flags().BoolS("i", "i", false, "Print contents of the InfoDictionary")
	rootCmd.Flags().Bool("leaks", false, "Run leaks at the end")
	rootCmd.Flags().BoolS("r", "r", false, "Real time")
	rootCmd.Flags().StringS("u", "u", "", "Find and print a property or user data property")
	rootCmd.Flags().Bool("warnings", false, "Print warnings")
	rootCmd.Flags().BoolS("x", "x", false, "Print output in xml format")
}
