package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var privateMACCmd = &cobra.Command{
	Use:   "privateMAC",
	Short: "set private MAC address mode (0/1)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(privateMACCmd).Standalone()
	carapace.Gen(privateMACCmd).PositionalCompletion(
		carapace.ActionValues("0", "1"),
	)
	rootCmd.AddCommand(privateMACCmd)
}
