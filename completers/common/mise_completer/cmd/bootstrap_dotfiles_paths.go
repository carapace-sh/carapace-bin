package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_pathsCmd = &cobra.Command{
	Use:   "paths",
	Short: "List tracked paths and how they are saved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_pathsCmd).Standalone()

	bootstrap_dotfiles_pathsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_pathsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_dotfiles_pathsCmd.Flags().Bool("noisy", false, "List the paths the watcher found changing constantly and saves less often")
	bootstrap_dotfiles_pathsCmd.Flags().String("preview", "", "List the files that tracking this path would save, without tracking it")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_pathsCmd)
}
