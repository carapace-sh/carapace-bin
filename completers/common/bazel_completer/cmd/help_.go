package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_Cmd = &cobra.Command{
	Use:   "",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_Cmd).Standalone()

	helpCmd.AddCommand(help_Cmd)
}
