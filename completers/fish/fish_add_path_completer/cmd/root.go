package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fish_add_path",
	Short: "Add directories to PATH",
	Long:  "https://fishshell.com/docs/current/cmds/fish_add_path.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "add to end")
	rootCmd.Flags().Bool("append", false, "add to end")
	rootCmd.Flags().BoolS("g", "g", false, "use global fish_user_paths")
	rootCmd.Flags().Bool("global", false, "use global fish_user_paths")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("m", "m", false, "move already-included directories")
	rootCmd.Flags().Bool("move", false, "move already-included directories")
	rootCmd.Flags().BoolS("n", "n", false, "dry run")
	rootCmd.Flags().Bool("dry-run", false, "dry run")
	rootCmd.Flags().BoolS("P", "P", false, "manipulate PATH directly")
	rootCmd.Flags().Bool("path", false, "manipulate PATH directly")
	rootCmd.Flags().BoolS("p", "p", false, "add to front")
	rootCmd.Flags().Bool("prepend", false, "add to front")
	rootCmd.Flags().BoolS("U", "U", false, "use universal fish_user_paths")
	rootCmd.Flags().Bool("universal", false, "use universal fish_user_paths")
	rootCmd.Flags().BoolS("v", "v", false, "verbose")
	rootCmd.Flags().Bool("verbose", false, "verbose")
}
