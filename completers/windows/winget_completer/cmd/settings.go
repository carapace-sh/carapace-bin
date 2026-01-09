package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:     "settings",
	Aliases: []string{"config"},
	Short:   "Open settings or set administrator settings",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settingsCmd).Standalone()

	rootCmd.AddCommand(settingsCmd)
}
