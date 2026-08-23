package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var authrestartCmd = &cobra.Command{
	Use:   "authrestart",
	Short: "Restart computer and unlock FileVault one time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authrestartCmd).Standalone()
	rootCmd.AddCommand(authrestartCmd)
}
