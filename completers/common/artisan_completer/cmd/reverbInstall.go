package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reverbInstallCmd = &cobra.Command{
	Use:   "reverb:install",
	Short: "Install the Reverb dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reverbInstallCmd).Standalone()

	rootCmd.AddCommand(reverbInstallCmd)
}
