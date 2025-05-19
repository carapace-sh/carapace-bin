package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uiAuthCmd = &cobra.Command{
	Use:   "ui:auth [--views] [--force] [--] [<type>]",
	Short: "Scaffold basic login and registration views and routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uiAuthCmd).Standalone()

	uiAuthCmd.Flags().Bool("force", false, "Overwrite existing views by default")
	uiAuthCmd.Flags().Bool("views", false, "Only scaffold the authentication views")
	rootCmd.AddCommand(uiAuthCmd)
}
