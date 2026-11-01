package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Inspect and preview aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_aliasCmd).Standalone()

	config_aliasCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_aliasCmd)
}
