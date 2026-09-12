package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:     "env",
	Short:   "Exports env vars to activate mise a single time",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	envCmd.Flags().StringP("shell", "s", "", "Shell type to generate env vars for")
	rootCmd.AddCommand(envCmd)

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{
		"shell": action.ActionShells(),
	})
}
