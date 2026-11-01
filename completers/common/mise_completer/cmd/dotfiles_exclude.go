package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_excludeCmd = &cobra.Command{
	Use:   "exclude",
	Short: "Never save paths matching a glob",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_excludeCmd).Standalone()

	dotfiles_excludeCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfilesCmd.AddCommand(dotfiles_excludeCmd)
}
