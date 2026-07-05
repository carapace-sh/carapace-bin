package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var disableJournalingCmd = &cobra.Command{
	Use:   "disableJournaling",
	Short: "Disable journaling on an HFS+ volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableJournalingCmd).Standalone()
	rootCmd.AddCommand(disableJournalingCmd)
	carapace.Gen(disableJournalingCmd).PositionalCompletion(fs.ActionBlockDevices())
}
