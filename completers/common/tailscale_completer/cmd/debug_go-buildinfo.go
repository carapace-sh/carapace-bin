package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_go_buildinfoCmd = &cobra.Command{
	Use:   "go-buildinfo",
	Short: "Print Go's runtime/debug.BuildInfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_go_buildinfoCmd).Standalone()

	debugCmd.AddCommand(debug_go_buildinfoCmd)
}
