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

	rootCmd.Flags().StringP("command", "o", "", "Execute a debugger command")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("source", "s", "", "Execute commands from a file")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
