package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var direnv_execCmd = &cobra.Command{
	Use:    "exec",
	Short:  "[internal] Print the environment for direnv to load",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(direnv_execCmd).Standalone()

	direnv_execCmd.Flags().BoolP("help", "h", false, "Print help")
	direnvCmd.AddCommand(direnv_execCmd)
}
