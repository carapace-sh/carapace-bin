package zig

import (
	"github.com/spf13/cobra"
)

func newHelpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "help",
		Short: "Print this help and exit",
	}

	return cmd
}
