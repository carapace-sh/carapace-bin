package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upnCmd = &cobra.Command{
	Use:   "upn",
	Short: "display user principal name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upnCmd).Standalone()
	rootCmd.AddCommand(upnCmd)
}
