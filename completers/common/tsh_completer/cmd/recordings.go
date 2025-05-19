package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recordingsCmd = &cobra.Command{
	Use:     "recordings",
	Short:   "View and control session recordings.",
	Aliases: []string{"recording"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordingsCmd).Standalone()

	rootCmd.AddCommand(recordingsCmd)
}
