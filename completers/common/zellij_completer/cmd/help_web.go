package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web server to serve terminal sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_webCmd).Standalone()

	helpCmd.AddCommand(help_webCmd)
}
