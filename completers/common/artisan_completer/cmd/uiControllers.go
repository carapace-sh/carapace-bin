package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uiControllersCmd = &cobra.Command{
	Use:   "ui:controllers",
	Short: "Scaffold the authentication controllers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uiControllersCmd).Standalone()

	rootCmd.AddCommand(uiControllersCmd)
}
