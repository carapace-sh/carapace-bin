package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debuguipromptCmd = &cobra.Command{
	Use:    "debuguiprompt",
	Short:  "show plain prompt",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuguipromptCmd).Standalone()

	debuguipromptCmd.Flags().StringP("prompt", "p", "", "prompt text")
	rootCmd.AddCommand(debuguipromptCmd)
}
