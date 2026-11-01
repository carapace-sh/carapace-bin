package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Start deploying a file, or copy its edits back to the source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_addCmd).Standalone()

	dotfiles_addCmd.Flags().Bool("changed", false, "Update the sources of every changed file managed in copy mode")
	dotfiles_addCmd.Flags().BoolP("dry-run", "n", false, "Print the config and source changes without writing anything")
	dotfiles_addCmd.Flags().BoolP("force", "f", false, "Overwrite existing sources without prompting")
	dotfiles_addCmd.Flags().BoolP("global", "g", false, "Write to the global config (the default)")
	dotfiles_addCmd.Flags().String("group", "", "Capture into the tree of this dotfile group")
	dotfiles_addCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_addCmd.Flags().BoolP("local", "l", false, "Write to the project config instead of the global config")
	dotfiles_addCmd.Flags().StringP("mode", "m", "", "Mode for a target that is not managed yet: symlink, symlink-each, copy, or template (default: the `dotfiles.default_mode` setting)")
	dotfiles_addCmd.Flags().Bool("no-apply", false, "Add the entry without applying it")
	dotfiles_addCmd.Flags().StringP("path", "p", "", "Write to this config file or directory")
	dotfiles_addCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	dotfiles_addCmd.Flags().StringP("source", "s", "", "Source path to use for a single target")
	dotfiles_addCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	dotfilesCmd.AddCommand(dotfiles_addCmd)
}
