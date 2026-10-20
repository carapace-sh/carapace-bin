package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var direnv_envrcCmd = &cobra.Command{
	Use:    "envrc",
	Short:  "[internal] Write an envrc file for direnv to consume",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(direnv_envrcCmd).Standalone()

	direnv_envrcCmd.Flags().BoolP("help", "h", false, "Print help")
	direnvCmd.AddCommand(direnv_envrcCmd)
}
