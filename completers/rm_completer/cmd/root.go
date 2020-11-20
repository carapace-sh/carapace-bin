package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rm",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("I", "I", false, "prompt once before removing more than three files, or when removing recursively")
	rootCmd.Flags().BoolP("dir", "d", false, "remove empty directories")
	rootCmd.Flags().BoolP("force", "f", false, "ignore nonexistent files and arguments, never prompt")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolS("i", "i", false, "prompt before every removal")
	rootCmd.Flags().String("interactive", "", "prompt according to WHEN: never, once (-I), or always (-i); without WHEN, prompt always")
	rootCmd.Flags().Bool("no-preserve-root", false, "do not treat '/' specially")
	rootCmd.Flags().Bool("one-file-system", false, "when removing a hierarchy recursively, skip any directory that is on a file system different from that of the corresponding command line argument")
	rootCmd.Flags().String("preserve-root", "", "do not remove '/' (default); with 'all', reject any command line argument on a separate device from its parent")
	rootCmd.Flags().BoolP("verbose", "v", false, "explain what is being done")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	rootCmd.Flag("interactive").NoOptDefVal = " "
	rootCmd.Flag("preserve-root").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"interactive": carapace.ActionValuesDescribed(
			"always", "prompt before every removal",
			"once", "prompt when removing many files",
		),
		"preserve-root": carapace.ActionValues("all"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(""),
	)
}
