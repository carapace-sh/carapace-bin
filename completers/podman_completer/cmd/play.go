package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:    "play",
	Short:  "Play containers, pods or volumes from a structured file",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(playCmd).Standalone()

	rootCmd.AddCommand(playCmd)
}
