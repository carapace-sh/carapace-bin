package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate a mise.toml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_configCmd).Standalone()

	generate_configCmd.Flags().BoolP("dry-run", "n", false, "Print the result instead of writing it to the file")
	generate_configCmd.Flags().BoolP("force", "f", false, "Replace an existing file with the starter template when not opening the editor")
	generate_configCmd.Flags().BoolP("global", "g", false, "Generate the global config file (~/.config/mise/config.toml)")
	generate_configCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_configCmd.Flags().StringP("tool-versions", "t", "", "Copy the tools from this .tool-versions file into the config, without opening the editor")
	generateCmd.AddCommand(generate_configCmd)

	carapace.Gen(generate_configCmd).FlagCompletion(carapace.ActionMap{
		"tool-versions": carapace.ActionFiles(),
	})

	carapace.Gen(generate_configCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
