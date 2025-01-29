package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_infoCmd = &cobra.Command{
	Use:   "info [options]",
	Short: "Display podman system information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_infoCmd).Standalone()

	system_infoCmd.Flags().BoolP("debug", "D", false, "Display additional debug information")
	system_infoCmd.Flags().StringP("format", "f", "", "Change the output format to JSON or a Go template")
	system_infoCmd.Flag("debug").Hidden = true
	systemCmd.AddCommand(system_infoCmd)
}
