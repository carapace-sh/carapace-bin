package zig

import (
	"github.com/spf13/cobra"
)

func newEnvCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Print lib path, std path, cache directory, and version",
	}

	return cmd
}
