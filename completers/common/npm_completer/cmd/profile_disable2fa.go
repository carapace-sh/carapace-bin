package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_disable2faCmd = &cobra.Command{
	Use:     "disable-2fa",
	Short:   "disable two-factor authorization",
	Aliases: []string{"disable-tfa", "disable2fa", "disabletfa"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_disable2faCmd).Standalone()

	profileCmd.AddCommand(profile_disable2faCmd)
}
