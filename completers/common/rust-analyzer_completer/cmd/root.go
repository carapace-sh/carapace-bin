package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rust-analyzer",
	Short: "LSP server for the Rust programming language",
	Long:  "https://github.com/rust-lang/rust-analyzer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().String("log-file", "", "Log to the specified file instead of stderr")
	rootCmd.Flags().Bool("no-log-buffering", false, "Flush log records to the file immediately")
	rootCmd.Flags().Bool("print-config-schema", false, "Dump a LSP config JSON schema")
	rootCmd.Flags().BoolP("quiet", "q", false, "Verbosity level")
	rootCmd.Flags().CountP("verbose", "v", "Verbosity level, can be repeated multiple times")
	rootCmd.Flags().Bool("version", false, "Print version")
	rootCmd.Flags().Bool("wait-dbg", false, "Wait until a debugger is attached to (requires debug build)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-file": carapace.ActionFiles(),
	})
}
