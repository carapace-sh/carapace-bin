package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addressCmd = &cobra.Command{
	Use:     "address",
	Aliases: []string{"addr", "a"},
	Short:   "protocol (IP or IPv6) address on a device",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addressCmd).Standalone()

	rootCmd.AddCommand(addressCmd)
}
