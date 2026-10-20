package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hookNotFoundCmd = &cobra.Command{
	Use:    "hook-not-found",
	Short:  "[internal] Handle a command the shell could not find",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookNotFoundCmd).Standalone()

	hookNotFoundCmd.Flags().BoolP("help", "h", false, "Print help")
	hookNotFoundCmd.Flags().StringP("shell", "s", "", "Shell type to generate script for")
	rootCmd.AddCommand(hookNotFoundCmd)

	carapace.Gen(hookNotFoundCmd).FlagCompletion(carapace.ActionMap{
		"shell": action.ActionShells(),
	})
}
