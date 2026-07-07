package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themeShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show theme(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themeShowCmd).Standalone()

	themeCmd.AddCommand(themeShowCmd)
}
