package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initLibCmd = &cobra.Command{
	Use:   "init-lib",
	Short: "Initialize a `zig build` library in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initLibCmd).Standalone()

	rootCmd.AddCommand(initLibCmd)
}
