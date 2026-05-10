package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newLibcCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "libc [paths_file]",
		Short: "Display native libc paths file or validate one",
	}

	cmd.Flags().StringP("target", "target", "", "The CPU architecture, OS, and ABI to detect libc for")

	carapace.Gen(cmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	return cmd
}
