package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var smbios11Cmd = &cobra.Command{
	Use:   "smbios11",
	Short: "List strings passed via SMBIOS Type #11",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(smbios11Cmd).Standalone()

	rootCmd.AddCommand(smbios11Cmd)
}
