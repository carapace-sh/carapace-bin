package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newRanlibCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ranlib [options]",
		Short: "Use Zig as a drop-in ranlib",
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(".a"),
	)

	return cmd
}
