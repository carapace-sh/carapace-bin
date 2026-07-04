package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lldb",
	Short: "LLVM debugger",
	Long:  "https://keith.github.io/xcode-manpages/lldb.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().StringS("Command|-o", "Command|-o", "", "")
	rootCmd.Flags().BoolS("Help|-h", "Help|-h", false, "")
	rootCmd.Flags().StringS("Source|-s", "Source|-s", "", "")
	rootCmd.Flags().BoolS("Version|-v", "Version|-v", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
