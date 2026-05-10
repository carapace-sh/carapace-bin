package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [options] [file] [-- [args]]",
		Short: "Create executable and run immediately",
	}

	addCompileFlags(cmd)
	addLinkFlags(cmd)

	carapace.Gen(cmd).PositionalCompletion(
		carapace.ActionFiles(".zig"),
	)

	return cmd
}
