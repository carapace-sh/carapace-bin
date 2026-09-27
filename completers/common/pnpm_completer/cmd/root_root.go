package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var root_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Print the effective `node_modules` directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(root_rootCmd).Standalone()

	root_rootCmd.Flags().BoolP("global", "g", false, "Print the global packages directory")
	root_rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(root_rootCmd)
}
