// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Return files to the version a checkpoint holds",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_rollbackCmd).Standalone()

	bootstrap_dotfiles_rollbackCmd.Flags().Bool("all", false, "With `--to`, restore everything the checkpoint covers")
	bootstrap_dotfiles_rollbackCmd.Flags().BoolP("dry-run", "n", false, "Show the plan without changing anything")
	bootstrap_dotfiles_rollbackCmd.Flags().Bool("force", false, "Replace a path whose type changed (file, symlink, directory)")
	bootstrap_dotfiles_rollbackCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_rollbackCmd.Flags().String("to", "", "The checkpoint to roll back to: numeric ID, `latest`, `latest~N`, or `commit:<sha>`")
	bootstrap_dotfiles_rollbackCmd.Flags().BoolP("yes", "y", false, "Apply without prompting")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_rollbackCmd)
}
