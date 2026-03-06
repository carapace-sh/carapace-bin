package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show configuration files & locations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_showCmd).Standalone()

	config_showCmd.Flags().Bool("full", false, "Run diagnostic checks (CI tools, commit generation, version)")
	config_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_showCmd)
}
