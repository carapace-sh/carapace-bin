package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var microfrontendsCmd = &cobra.Command{
	Use:     "microfrontends",
	Aliases: []string{"mf"},
	Short:   "Manage microfrontends groups",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(microfrontendsCmd).Standalone()

	rootCmd.AddCommand(microfrontendsCmd)
}
