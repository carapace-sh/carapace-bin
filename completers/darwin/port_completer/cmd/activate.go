package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate the installed port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(activateCmd).Standalone()
	rootCmd.AddCommand(activateCmd)
}
