package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newCcCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cc [options]",
		Short: "Use Zig as a drop-in C compiler",
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(".c", ".h", ".o", ".s", ".S"),
	)

	return cmd
}
