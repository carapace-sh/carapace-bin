package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_forgeCmd = &cobra.Command{
	Use:   "forge",
	Short: "View and manage forge configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_forgeCmd).Standalone()

	config_forgeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_forgeCmd)
}
