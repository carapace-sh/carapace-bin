package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fileproviderctl",
	Short: "introspect file provider extensions",
	Long:  "https://keith.github.io/xcode-manpages/fileproviderctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"enumerate", "Run an interactive enumeration of the specified provider",
			"ls", "Run an interactive enumeration of the specified provider",
			"materialize", "Cause the specified item to be written on disk",
			"validate", "Run the validation suite against the specified provider",
			"dump", "Dump the state of the file provider subsystem",
		),
	)
}