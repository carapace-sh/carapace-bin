package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fortifyInstallCmd = &cobra.Command{
	Use:   "fortify:install",
	Short: "Install all of the Fortify resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fortifyInstallCmd).Standalone()

	rootCmd.AddCommand(fortifyInstallCmd)
}
