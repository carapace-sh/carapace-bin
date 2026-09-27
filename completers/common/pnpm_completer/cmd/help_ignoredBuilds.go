package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ignoredBuildsCmd = &cobra.Command{
	Use:   "ignored-builds",
	Short: "Print the list of packages with blocked build scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ignoredBuildsCmd).Standalone()

	helpCmd.AddCommand(help_ignoredBuildsCmd)
}
