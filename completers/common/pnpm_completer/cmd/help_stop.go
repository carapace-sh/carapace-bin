package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Runs a package's \"stop\" script, if one was provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stopCmd).Standalone()

	helpCmd.AddCommand(help_stopCmd)
}
