package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [tool@version]...",
	Short: "Add tool versions to mise config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()
	useCmd.Flags().BoolP("global", "g", false, "Use the global config")
	useCmd.Flags().BoolP("path", "p", false, "Use .mise.toml in the current directory")
	useCmd.Flags().Bool("pin", false, "Save exact versions")
	useCmd.Flags().Bool("fuzzy", false, "Save fuzzy versions")
	rootCmd.AddCommand(useCmd)
}
