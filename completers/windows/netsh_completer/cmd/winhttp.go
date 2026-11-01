package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var winhttpCmd = &cobra.Command{
	Use:   "winhttp",
	Short: "WinHTTP proxy configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(winhttpCmd).Standalone()
	rootCmd.AddCommand(winhttpCmd)
}
