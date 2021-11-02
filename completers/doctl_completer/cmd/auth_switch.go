package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var auth_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switches between authentication contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_switchCmd).Standalone()
	authCmd.AddCommand(auth_switchCmd)
}
