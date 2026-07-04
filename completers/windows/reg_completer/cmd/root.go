package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reg",
	Short: "perform operations on registry subkey information and values",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/reg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "add a new subkey or entry to the registry",
			"compare", "compare registry subkeys or entries",
			"copy", "copy a registry entry",
			"delete", "delete a subkey or entries",
			"export", "export subkeys and values to a file",
			"import", "copy a file into the registry",
			"load", "load a subkey into a different subkey",
			"query", "display subkeys and entries",
			"restore", "write saved subkeys back to the registry",
			"save", "save a subkey to a file",
			"unload", "unload a subkey",
		),
	)
}
