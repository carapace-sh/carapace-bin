package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newStripCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "strip [options]",
		Short: "Use Zig as a drop-in strip",
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(".o", ".elf", ".so", ".a"),
	)

	return cmd
}
