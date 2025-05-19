package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui [--auth] [--option [OPTION]] [--] <type>",
	Short: "Swap the front-end scaffolding for the application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uiCmd).Standalone()

	uiCmd.Flags().Bool("auth", false, "Install authentication UI scaffolding")
	uiCmd.Flags().String("option", "", "Pass an option to the preset command")
	rootCmd.AddCommand(uiCmd)
}
