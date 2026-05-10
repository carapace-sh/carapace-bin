package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newTranslateCCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "translate-c [options] [file]",
		Short: "Convert C code to Zig code",
	}

	addCompileFlags(cmd)

	carapace.Gen(cmd).PositionalCompletion(
		carapace.ActionFiles(".c", ".h"),
	)

	return cmd
}
