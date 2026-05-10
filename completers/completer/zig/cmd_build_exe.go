package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newBuildExeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build-exe [options] [files]",
		Short: "Create executable from source or object files",
	}

	addCompileFlags(cmd)
	addLinkFlags(cmd)

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig", ".o", ".c", ".cpp", ".cxx", ".cc", ".s", ".S"),
	)

	return cmd
}
