package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reloadAtfCmd = &cobra.Command{
	Use:   "reload-atf",
	Short: "Reloads the active trial factors",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(reloadAtfCmd).Standalone()
	rootCmd.AddCommand(reloadAtfCmd)
}
