package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chmod",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
    carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("changes", "c", false, "like verbose but report only when a change is made")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("no-preserve-root", false, "do not treat '/' specially (the default)")
	rootCmd.Flags().Bool("preserve-root", false, "fail to operate recursively on '/'")
	rootCmd.Flags().BoolP("recursive", "R", false, "change files and directories recursively")
	rootCmd.Flags().String("reference", "", "use RFILE's mode instead of MODE values")
	rootCmd.Flags().StringP("silent,", "f", "", "suppress most error messages")
	rootCmd.Flags().BoolP("verbose", "v", false, "output a diagnostic for every file processed")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

    carapace.Gen(rootCmd).PositionalCompletion(
        carapace.ActionValues(""),
    )

    carapace.Gen(rootCmd).PositionalAnyCompletion(
        carapace.ActionFiles(""),
    )
}
