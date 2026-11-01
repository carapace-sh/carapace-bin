package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "enable or disable session port mapping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portCmd).Standalone()
	rootCmd.AddCommand(portCmd)
}
