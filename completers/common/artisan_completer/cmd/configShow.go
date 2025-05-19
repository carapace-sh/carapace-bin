package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configShowCmd = &cobra.Command{
	Use:   "config:show <config>",
	Short: "Display all of the values for a given configuration file or key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configShowCmd).Standalone()

	rootCmd.AddCommand(configShowCmd)
}
