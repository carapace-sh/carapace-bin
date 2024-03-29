package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var advertiseCmd = &cobra.Command{
	Use:   "advertise [on|off|broadcast|peripheral]",
	Short: "Enable/disable advertising",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advertiseCmd).Standalone()
	rootCmd.AddCommand(advertiseCmd)
	carapace.Gen(advertiseCmd).PositionalCompletion(
		carapace.ActionValues("on", "off", "broadcast", "peripheral").StyleF(style.ForKeyword),
	)
}
