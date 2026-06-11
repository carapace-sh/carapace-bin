package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initExeCmd = &cobra.Command{
	Use:   "init-exe",
	Short: "Initialize a `zig build` application in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initExeCmd).Standalone()

	rootCmd.AddCommand(initExeCmd)
}
