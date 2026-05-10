package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newBuildObjCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build-obj [options] [files]",
		Short: "Create object from source or object files",
	}

	addCompileFlags(cmd)

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig", ".o", ".c", ".cpp", ".cxx", ".cc", ".s", ".S"),
	)

	return cmd
}
