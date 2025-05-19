package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaInstallCmd = &cobra.Command{
	Use:   "nova:install",
	Short: "Install all of the Nova resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaInstallCmd).Standalone()

	rootCmd.AddCommand(novaInstallCmd)
}
