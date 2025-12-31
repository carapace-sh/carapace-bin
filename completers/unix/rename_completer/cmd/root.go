package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename files",
	Long:  "https://www.man7.org/linux/man-pages/man1/rename.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "replace all occurrences")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("interactive", "i", false, "prompt before overwrite")
	rootCmd.Flags().BoolP("last", "l", false, "replace only the last occurrence")
	rootCmd.Flags().BoolP("no-act", "n", false, "do not make any changes")
	rootCmd.Flags().BoolP("no-overwrite", "o", false, "don't overwrite existing files")
	rootCmd.Flags().BoolP("symlink", "s", false, "act on the target of symlinks")
	rootCmd.Flags().BoolP("verbose", "v", false, "explain what is being done")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
	)
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
