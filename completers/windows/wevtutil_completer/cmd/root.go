package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wevtutil",
	Short: "retrieve information about event logs and publishers",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/wevtutil",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"el", "enumerate event log names",
			"gl", "get event log configuration",
			"sl", "set event log configuration",
			"ep", "enumerate event publishers",
			"gp", "get event publisher metadata",
			"im", "install event publisher manifest",
			"um", "uninstall event publisher manifest",
			"qe", "query events from an event log",
			"ali", "get log names and aliases",
			"epl", "export events to a file",
			"ipl", "import events from a file",
			"cl", "clear an event log",
		),
	)
}
