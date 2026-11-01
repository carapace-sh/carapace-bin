package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var winsockCmd = &cobra.Command{
	Use:   "winsock",
	Short: "Winsock configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(winsockCmd).Standalone()
	rootCmd.AddCommand(winsockCmd)
}
