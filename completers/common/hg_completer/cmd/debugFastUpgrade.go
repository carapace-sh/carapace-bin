package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugFastUpgradeCmd = &cobra.Command{
	Use:     "debug::fast-upgrade",
	Short:   "",
	Aliases: []string{"debug::fastupgrade"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugFastUpgradeCmd).Standalone()

	rootCmd.AddCommand(debugFastUpgradeCmd)
}
