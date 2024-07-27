package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telescopeResumeCmd = &cobra.Command{
	Use:   "telescope:resume",
	Short: "Unpause all Telescope watchers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telescopeResumeCmd).Standalone()

	rootCmd.AddCommand(telescopeResumeCmd)
}
