package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scripts_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_listCmd).Standalone()

	scripts_listCmd.Flags().BoolP("help", "h", false, "Print help")
	scriptsCmd.AddCommand(scripts_listCmd)
}
