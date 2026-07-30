package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nslookup",
	Short: "display information that you can use to diagnose DNS infrastructure",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/nslookup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
