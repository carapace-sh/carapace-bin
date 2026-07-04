package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpkgCmd = &cobra.Command{
	Use:   "mpkg",
	Short: "Create a macOS installer metapackage of a port and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpkgCmd).Standalone()
	rootCmd.AddCommand(mpkgCmd)
}
