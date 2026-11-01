package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysprep",
	Short: "generalize a Windows installation for deployment",
	Long:  "https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/sysprep-command-line-options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
