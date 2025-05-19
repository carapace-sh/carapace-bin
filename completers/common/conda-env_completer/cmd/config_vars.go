package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "Interact with environment variables associated with Conda environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_varsCmd).Standalone()

	configCmd.AddCommand(config_varsCmd)
}
