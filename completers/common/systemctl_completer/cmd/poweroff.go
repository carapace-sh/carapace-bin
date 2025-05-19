package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var poweroffCmd = &cobra.Command{
	Use:     "poweroff",
	Short:   "Shut down and power-off the system",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(poweroffCmd).Standalone()

	rootCmd.AddCommand(poweroffCmd)
}
