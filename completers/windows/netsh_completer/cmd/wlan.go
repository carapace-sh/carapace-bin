package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wlanCmd = &cobra.Command{
	Use:   "wlan",
	Short: "wireless LAN configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wlanCmd).Standalone()
	rootCmd.AddCommand(wlanCmd)
}
