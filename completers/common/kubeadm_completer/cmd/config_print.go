package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_printCmd).Standalone()

	configCmd.AddCommand(config_printCmd)
}
