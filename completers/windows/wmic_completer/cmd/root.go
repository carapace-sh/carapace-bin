package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wmic",
	Short: "display WMI information in an interactive command shell",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/wmic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"baseboard", "baseboard management",
			"bios", "BIOS management",
			"computersystem", "computer system management",
			"cpu", "CPU management",
			"desktop", "desktop management",
			"diskdrive", "disk drive management",
			"logicaldisk", "logical disk management",
			"memorychip", "memory chip management",
			"nicconfig", "network adapter configuration",
			"os", "operating system management",
			"process", "process management",
			"product", "installed products",
			"service", "service management",
			"sounddev", "sound device management",
			"useraccount", "user account management",
			"volume", "volume management",
		),
	)
}
