package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reverbRestartCmd = &cobra.Command{
	Use:   "reverb:restart",
	Short: "Restart the Reverb server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reverbRestartCmd).Standalone()

	rootCmd.AddCommand(reverbRestartCmd)
}
