package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var app_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "set default application policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(app_defaultCmd)

	appCmd.AddCommand(app_defaultCmd)
}
