package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:   "power [on|off]",
	Short: "Set controller power state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(powerCmd).Standalone()
	rootCmd.AddCommand(powerCmd)

	carapace.Gen(powerCmd).PositionalCompletion(
		carapace.ActionValues("on", "off").StyleF(style.ForKeyword),
	)
}
