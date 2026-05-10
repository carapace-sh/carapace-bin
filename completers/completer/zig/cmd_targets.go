package zig

import (
	"github.com/spf13/cobra"
)

func newTargetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "targets",
		Short: "List available compilation targets",
	}

	return cmd
}
