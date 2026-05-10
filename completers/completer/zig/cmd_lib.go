package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newLibCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lib [options]",
		Short: "Use Zig as a drop-in lib.exe",
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	return cmd
}
