package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new, empty repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("bare", false, "Create a bare repository (no working tree)")
	initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	initCmd.Flags().StringP("initial-branch", "b", "", "Initial branch name (overrides `init.defaultBranch` from config)")
	initCmd.Flags().String("object-format", "", "Object hash algorithm for the new repository (default SHA-1; omitted on reinit preserves format)")
	initCmd.Flags().String("ref-format", "files", "Ref storage backend for the new repository")
	initCmd.Flags().String("separate-git-dir", "", "Place the git directory here and link from the work tree via a gitfile")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"separate-git-dir": carapace.ActionFiles(),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
