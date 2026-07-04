package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "certutil",
	Short: "command-line program to manage certificates and certificate services",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/certutil",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
