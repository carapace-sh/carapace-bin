package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information for an application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	infoCmd.Flags().String("app", "", "Application specifier")
	infoCmd.Flags().Bool("long", false, "Long format")
	infoCmd.Flags().String("only", "", "Only show specific information item key")
	rootCmd.AddCommand(infoCmd)
}
