package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var app_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list application profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(app_listCmd)

	appCmd.AddCommand(app_listCmd)
}
