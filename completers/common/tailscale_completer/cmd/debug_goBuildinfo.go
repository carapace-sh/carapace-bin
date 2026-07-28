package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_goBuildinfoCmd = &cobra.Command{
	Use:   "go-buildinfo",
	Short: "Print Go's runtime/debug.BuildInfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_goBuildinfoCmd).Standalone()

	debugCmd.AddCommand(debug_goBuildinfoCmd)
}
