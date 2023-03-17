package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var alpha_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Inspect authorization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_authCmd).Standalone()

	alphaCmd.AddCommand(alpha_authCmd)
}
