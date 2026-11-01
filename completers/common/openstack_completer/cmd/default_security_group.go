package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var default_security_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(default_security_groupCmd).Standalone()

	default_securityCmd.AddCommand(default_security_groupCmd)
}
