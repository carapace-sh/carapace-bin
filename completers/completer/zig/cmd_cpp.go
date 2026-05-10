package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newCppCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "c++ [options]",
		Short: "Use Zig as a drop-in C++ compiler",
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(".cpp", ".cxx", ".cc", ".h", ".hpp", ".o", ".s", ".S"),
	)

	return cmd
}
