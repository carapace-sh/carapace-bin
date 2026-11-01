package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logonCmd = &cobra.Command{
	Use:   "logon",
	Short: "enable or disable logon from sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logonCmd).Standalone()
	rootCmd.AddCommand(logonCmd)
}
