package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "powercfg",
	Short: "control power settings and configure power plans",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/powercfg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"list", "list all available power plans",
			"query", "display contents of a power plan",
			"change", "modify a setting in the current power plan",
			"setactive", "make a power plan active",
			"getactivescheme", "retrieve the active power plan",
			"delete", "delete a power plan",
			"duplicatescheme", "duplicate a power plan",
			"import", "import a power plan from a file",
			"export", "export a power plan to a file",
			"aliases", "display aliases and corresponding identifiers",
			"devicequery", "query devices that meet certain criteria",
			"wakesources", "list devices that wake the system",
			"lastwake", "report information about what woke the system",
		),
	)
}
