package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Ask launchservicesd to restart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()
	rootCmd.AddCommand(restartCmd)
}
