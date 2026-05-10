package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test [options] [file]",
		Short: "Perform unit testing",
	}

	addCompileFlags(cmd)
	addLinkFlags(cmd)

	cmd.Flags().String("test-filter", "", "Skip tests that do not match any filter")
	cmd.Flags().String("test-name-prefix", "", "Add prefix to all tests")
	cmd.Flags().String("test-cmd", "", "Specify test executor executable")
	cmd.Flags().Bool("test-cmd-bin", false, "Appends test binary path to test cmd args")
	cmd.Flags().Bool("test-no-exec", false, "Compiles test binary without running it")
	cmd.Flags().String("test-runner", "", "Specify a custom test runner")
	cmd.Flags().Bool("listen", false, "Enable the compiler to listen for incoming requests")
	cmd.Flags().String("listen", "", "Enable the compiler to listen for incoming requests on specified address")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"test-runner": carapace.ActionFiles(".zig"),
	})

	carapace.Gen(cmd).PositionalCompletion(
		carapace.ActionFiles(".zig"),
	)

	return cmd
}
