package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telescopePauseCmd = &cobra.Command{
	Use:   "telescope:pause",
	Short: "Pause all Telescope watchers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telescopePauseCmd).Standalone()

	rootCmd.AddCommand(telescopePauseCmd)
}
