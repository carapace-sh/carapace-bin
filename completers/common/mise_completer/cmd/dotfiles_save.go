package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a checkpoint of the tracked files now",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_saveCmd).Standalone()

	dotfiles_saveCmd.Flags().Bool("best-effort", false, "Warn instead of failing when the checkpoint cannot be written")
	dotfiles_saveCmd.Flags().StringP("description", "d", "", "A description for the checkpoint")
	dotfiles_saveCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_saveCmd.Flags().StringSlice("label", nil, "A label to find the checkpoint by later")
	dotfiles_saveCmd.Flags().Bool("re-encrypt", false, "Re-encrypt every saved encrypted file to the current recipients")
	dotfiles_saveCmd.Flags().String("task", "", "The ID of the task an agent is working on, recorded with the checkpoint")
	dotfiles_saveCmd.Flags().String("trigger", "save", "Record the checkpoint as a save, agent, or update checkpoint")
	dotfilesCmd.AddCommand(dotfiles_saveCmd)
}
