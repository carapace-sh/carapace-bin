package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Repair config files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_fixCmd).Standalone()
	configCmd.AddCommand(config_fixCmd)
}
