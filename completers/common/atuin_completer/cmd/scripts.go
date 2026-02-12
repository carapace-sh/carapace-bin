package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Manage your scripts with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scriptsCmd).Standalone()

	scriptsCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(scriptsCmd)
}
