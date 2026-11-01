package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_configCmd = &cobra.Command{
	Use:   "config <command> [flags]",
	Short: "Configure security scan profiles for a project. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_configCmd).Standalone()

	securityCmd.AddCommand(security_configCmd)
}
