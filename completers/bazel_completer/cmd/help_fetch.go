package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_fetchCmd).Standalone()

	helpCmd.AddCommand(help_fetchCmd)
}
