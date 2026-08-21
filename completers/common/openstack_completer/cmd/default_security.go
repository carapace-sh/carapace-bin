package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var default_securityCmd = &cobra.Command{
	Use:   "security",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(default_securityCmd).Standalone()

	defaultCmd.AddCommand(default_securityCmd)
}
