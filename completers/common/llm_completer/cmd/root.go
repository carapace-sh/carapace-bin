package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "llm",
	Short: "Access large language models from the command-line",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("version", false, "Show the version and exit.")
}

// applyFlagCompletion registers completion for flags that exist on the command
func applyFlagCompletion(cmd *cobra.Command, actions carapace.ActionMap) {
	available := make(carapace.ActionMap)
	for name, a := range actions {
		if cmd.LocalFlags().Lookup(name) != nil {
			available[name] = a
		}
	}
	carapace.Gen(cmd).FlagCompletion(available)
}
