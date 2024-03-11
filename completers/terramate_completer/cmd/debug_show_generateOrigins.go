package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_show_generateOriginsCmd = &cobra.Command{
	Use:   "generate-origins",
	Short: "Show generate debug information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_show_generateOriginsCmd).Standalone()

	debug_showCmd.AddCommand(debug_show_generateOriginsCmd)
}
