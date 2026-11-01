package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hostsAddCmd = &cobra.Command{
	Use:   "add LABEL ENDPOINT...",
	Short: "Add a host to your home",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostsAddCmd).Standalone()

	hostsAddCmd.Flags().String("icon", "", "show the icon NAME next to the host")

	hostsCmd.AddCommand(hostsAddCmd)
}
