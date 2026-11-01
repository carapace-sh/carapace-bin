package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Start deploying a file, or copy its edits back to the source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_addCmd).Standalone()

	bootstrap_dotfiles_addCmd.Flags().Bool("changed", false, "Update the sources of every changed file managed in copy mode")
	bootstrap_dotfiles_addCmd.Flags().BoolP("dry-run", "n", false, "Print the config and source changes without writing anything")
	bootstrap_dotfiles_addCmd.Flags().BoolP("force", "f", false, "Overwrite existing sources without prompting")
	bootstrap_dotfiles_addCmd.Flags().BoolP("global", "g", false, "Write to the global config (the default)")
	bootstrap_dotfiles_addCmd.Flags().String("group", "", "Capture into the tree of this dotfile group")
	bootstrap_dotfiles_addCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_addCmd.Flags().BoolP("local", "l", false, "Write to the project config instead of the global config")
	bootstrap_dotfiles_addCmd.Flags().StringP("mode", "m", "", "Mode for a target that is not managed yet: symlink, symlink-each, copy, or template (default: the `dotfiles.default_mode` setting)")
	bootstrap_dotfiles_addCmd.Flags().Bool("no-apply", false, "Add the entry without applying it")
	bootstrap_dotfiles_addCmd.Flags().StringP("path", "p", "", "Write to this config file or directory")
	bootstrap_dotfiles_addCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	bootstrap_dotfiles_addCmd.Flags().StringP("source", "s", "", "Source path to use for a single target")
	bootstrap_dotfiles_addCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_addCmd)
}
