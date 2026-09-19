package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugignoreCmd = &cobra.Command{
	Use:    "debugignore",
	Short:  "display the combined ignore pattern and information about ignored files",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugignoreCmd).Standalone()

	rootCmd.AddCommand(debugignoreCmd)
}
