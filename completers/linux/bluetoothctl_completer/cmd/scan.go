package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan <on|off|bredr|le>",
	Short: "Scan for devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scanCmd).Standalone()
	rootCmd.AddCommand(scanCmd)

	carapace.Gen(scanCmd).PositionalCompletion(
		carapace.ActionValues("on", "off", "bredr", "le").StyleF(style.ForKeyword),
	)
}
