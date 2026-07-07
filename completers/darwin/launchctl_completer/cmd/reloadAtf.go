package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reloadAtfCmd = &cobra.Command{
	Use:   "reload-atf",
	Short: "Reload the active trial factors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadAtfCmd).Standalone()
	rootCmd.AddCommand(reloadAtfCmd)
}
