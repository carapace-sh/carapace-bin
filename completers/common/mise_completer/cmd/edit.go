package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit mise.toml interactively",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	editCmd.Flags().BoolP("dry-run", "n", false, "Print the result instead of writing it to the file")
	editCmd.Flags().BoolP("force", "f", false, "Replace an existing file with the starter template when not opening the editor")
	editCmd.Flags().BoolP("global", "g", false, "Edit the global config file (~/.config/mise/config.toml)")
	editCmd.Flags().BoolP("help", "h", false, "Print help")
	editCmd.Flags().StringP("tool-versions", "t", "", "Copy the tools from this .tool-versions file into the config, without opening the editor")
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).FlagCompletion(carapace.ActionMap{
		"tool-versions": carapace.ActionFiles(),
	})

	carapace.Gen(editCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
