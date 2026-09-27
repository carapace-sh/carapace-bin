package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Safely remove `node_modules` directories from the current project (or every workspace project) without following NTFS junctions into their targets. A `clean` script in `package.json` overrides the built-in command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cleanCmd).Standalone()

	helpCmd.AddCommand(help_cleanCmd)
}
