package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addhostCmd = &cobra.Command{
	Use:   "addhost",
	Short: "add a host-to-realm mapping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addhostCmd).Standalone()
	rootCmd.AddCommand(addhostCmd)
}
