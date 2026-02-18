package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var root_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Show the current workspace root directory (shortcut for `jj workspace root`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(root_rootCmd).Standalone()

	root_rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(root_rootCmd)
}
