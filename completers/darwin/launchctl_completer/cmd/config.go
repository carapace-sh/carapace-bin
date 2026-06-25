package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modifies persistent configuration parameters for launchd domains",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()
	rootCmd.AddCommand(configCmd)
}
