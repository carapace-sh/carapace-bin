package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rustup",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Root()
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().BoolP("quiet", "q", false, "Disable progress output")
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")
}
