package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_config_statusCmd = &cobra.Command{
	Use:   "status <profile> [flags]",
	Short: "Show the status of a security scan profile for a project. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_config_statusCmd).Standalone()

	security_configCmd.AddCommand(security_config_statusCmd)
}
