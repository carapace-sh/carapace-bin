package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mv",
	Short: "move files",
	Long:  "https://keith.github.io/xcode-manpages/mv.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation before overwriting the destination path")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("interactive", "i", false, "Cause mv to write a prompt to the standard error output before moving a file")
	rootCmd.Flags().BoolP("no-clobber", "n", false, "Do not overwrite an existing file")
	rootCmd.Flags().BoolP("verbose", "v", false, "Cause mv to be verbose, showing files after they are moved")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
