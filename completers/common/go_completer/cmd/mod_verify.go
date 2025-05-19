package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify dependencies have expected content",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_verifyCmd).Standalone()
	mod_verifyCmd.Flags().SetInterspersed(false)

	modCmd.AddCommand(mod_verifyCmd)
}
