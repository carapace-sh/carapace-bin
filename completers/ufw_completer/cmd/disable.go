package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "disables the firewall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableCmd)

	rootCmd.AddCommand(disableCmd)
}
