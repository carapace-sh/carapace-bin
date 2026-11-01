package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debuggerCmd = &cobra.Command{
	Use:   "debugger",
	Short: "Run tests under an instruction-level debugger",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuggerCmd).Standalone()

	debuggerCmd.Flags().Bool("gdb", false, "Drive tests over sbpf's gdb-stub instead of reading dumped trace files")
	debuggerCmd.Flags().BoolP("help", "h", false, "Print help")
	debuggerCmd.Flags().Bool("skip-build", false, "Skip `cargo build-sbf`")
	debuggerCmd.Flags().Bool("skip-lint", false, "Forwarded to the underlying `anchor test` invocation")
	debuggerCmd.Flags().Bool("skip-run", false, "Skip the build+test phase and open the TUI over existing traces")
	rootCmd.AddCommand(debuggerCmd)
}
