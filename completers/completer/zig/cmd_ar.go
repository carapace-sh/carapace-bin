package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arCmd = &cobra.Command{
	Use:   "ar",
	Short: "Use Zig as a drop-in archiver",
}

func newArCmd() *cobra.Command {
	arCmd := &cobra.Command{
		Use:   "ar",
		Short: "Use Zig as a drop-in archiver",
	}
	carapace.Gen(arCmd).Standalone()
	return arCmd
}
