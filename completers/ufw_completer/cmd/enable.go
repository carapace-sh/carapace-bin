package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "enables the firewall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableCmd)

	rootCmd.AddCommand(enableCmd)
}
