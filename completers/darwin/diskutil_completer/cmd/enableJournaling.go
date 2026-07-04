package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var enableJournalingCmd = &cobra.Command{
	Use:   "enableJournaling",
	Short: "Enable journaling on an HFS+ volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableJournalingCmd).Standalone()
	rootCmd.AddCommand(enableJournalingCmd)
	carapace.Gen(enableJournalingCmd).PositionalCompletion(fs.ActionBlockDevices())
}
