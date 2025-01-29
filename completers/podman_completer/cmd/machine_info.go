package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_infoCmd = &cobra.Command{
	Use:   "info [options]",
	Short: "Display machine host info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_infoCmd).Standalone()

	machine_infoCmd.Flags().StringP("format", "f", "", "Change the output format to JSON or a Go template")
	machineCmd.AddCommand(machine_infoCmd)
}
