package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available themes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themeListCmd).Standalone()

	themeCmd.AddCommand(themeListCmd)
}
