package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a checkpoint of the tracked files now",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_saveCmd).Standalone()

	bootstrap_dotfiles_saveCmd.Flags().Bool("best-effort", false, "Warn instead of failing when the checkpoint cannot be written")
	bootstrap_dotfiles_saveCmd.Flags().StringP("description", "d", "", "A description for the checkpoint")
	bootstrap_dotfiles_saveCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_saveCmd.Flags().StringSlice("label", nil, "A label to find the checkpoint by later")
	bootstrap_dotfiles_saveCmd.Flags().Bool("re-encrypt", false, "Re-encrypt every saved encrypted file to the current recipients")
	bootstrap_dotfiles_saveCmd.Flags().String("task", "", "The ID of the task an agent is working on, recorded with the checkpoint")
	bootstrap_dotfiles_saveCmd.Flags().String("trigger", "save", "Record the checkpoint as a save, agent, or update checkpoint")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_saveCmd)
}
