package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repadmin",
	Short: "manage Active Directory replication between domain controllers",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/repadmin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
