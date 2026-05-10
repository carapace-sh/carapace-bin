package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newDlltoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dlltool [options]",
		Short: "Use Zig as a drop-in dlltool",
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	return cmd
}
