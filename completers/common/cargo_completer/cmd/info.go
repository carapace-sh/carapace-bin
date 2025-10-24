package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolP("help", "h", false, "Print help")
	infoCmd.Flags().String("index", "", "Registry index URL to search packages in")
	infoCmd.Flags().String("registry", "", "Registry to search packages in")
	rootCmd.AddCommand(infoCmd)

	// TODO completion
}
