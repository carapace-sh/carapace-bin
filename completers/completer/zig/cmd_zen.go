package zig

import (
	"github.com/spf13/cobra"
)

func newZenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zen",
		Short: "Print Zen of Zig and exit",
	}

	return cmd
}
