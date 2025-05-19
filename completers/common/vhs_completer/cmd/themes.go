package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themesCmd = &cobra.Command{
	Use:   "themes",
	Short: "List all the available themes, one per line",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themesCmd).Standalone()

	themesCmd.Flags().Bool("markdown", false, "output as markdown")
	themesCmd.Flag("markdown").Hidden = true
	rootCmd.AddCommand(themesCmd)
}
