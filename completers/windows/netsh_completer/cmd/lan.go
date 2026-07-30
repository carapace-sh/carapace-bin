package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lanCmd = &cobra.Command{
	Use:   "lan",
	Short: "wired LAN configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lanCmd).Standalone()
	rootCmd.AddCommand(lanCmd)
}
