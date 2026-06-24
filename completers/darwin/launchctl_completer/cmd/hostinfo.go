package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hostinfoCmd = &cobra.Command{
	Use:   "hostinfo",
	Short: "Prints port information about the host",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(hostinfoCmd).Standalone()
	rootCmd.AddCommand(hostinfoCmd)
}
