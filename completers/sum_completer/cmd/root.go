package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sum",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Root()
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolS("r", "r", false, "use BSD sum algorithm, use 1K blocks")
	rootCmd.Flags().BoolP("sysv", "s", false, "use System V sum algorithm, use 512 bytes blocks")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
