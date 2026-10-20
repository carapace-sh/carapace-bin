package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoMode_defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Print the default auto mode environment, allow, soft_deny, and hard_deny rules as JSON",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoMode_defaultsCmd).Standalone()

	autoMode_defaultsCmd.Flags().BoolP("help", "h", false, "Display help for command")
	autoMode_defaultsCmd.Flags().String("label", "", "Show only rules whose label starts with this prefix (case-insensitive)")
	autoModeCmd.AddCommand(autoMode_defaultsCmd)
}
