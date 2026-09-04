package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var msbuildCmd = &cobra.Command{
	Use:   "msbuild",
	Short: "run MSBuild",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(msbuildCmd).Standalone()
	rootCmd.AddCommand(msbuildCmd)
}
