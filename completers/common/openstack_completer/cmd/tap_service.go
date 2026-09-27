package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_serviceCmd).Standalone()

	tapCmd.AddCommand(tap_serviceCmd)
}
