package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Return files to the version a checkpoint holds",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_rollbackCmd).Standalone()

	dotfiles_rollbackCmd.Flags().Bool("all", false, "With `--to`, restore everything the checkpoint covers")
	dotfiles_rollbackCmd.Flags().BoolP("dry-run", "n", false, "Show the plan without changing anything")
	dotfiles_rollbackCmd.Flags().Bool("force", false, "Replace a path whose type changed (file, symlink, directory)")
	dotfiles_rollbackCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_rollbackCmd.Flags().String("to", "", "The checkpoint to roll back to: numeric ID, `latest`, `latest~N`, or `commit:<sha>`")
	dotfiles_rollbackCmd.Flags().BoolP("yes", "y", false, "Apply without prompting")
	dotfilesCmd.AddCommand(dotfiles_rollbackCmd)
}
