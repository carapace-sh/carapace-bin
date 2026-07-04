package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nvram",
	Short: "manipulate firmware variables",
	Long:  "https://keith.github.io/xcode-manpages/nvram.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("clear", "c", false, "Clear all firmware variables")
	rootCmd.Flags().BoolP("delete", "d", false, "Delete the specified variable")
	rootCmd.Flags().StringP("file", "f", "", "Use the specified file for input/output")
	rootCmd.Flags().BoolP("print", "p", false, "Print all firmware variables")
	rootCmd.Flags().BoolP("xml", "x", false, "Print in XML format")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
