package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Check for updates and install if available; package-manager installs update through the package…",
	Aliases: []string{"upgrade"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(updateCmd)
}
