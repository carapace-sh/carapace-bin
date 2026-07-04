package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dxdiag",
	Short: "display DirectX diagnostic tool information",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/dxdiag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("dontskip", false, "do not skip warnings")
	rootCmd.Flags().Bool("whql", false, "check WHQL signatures")
}
