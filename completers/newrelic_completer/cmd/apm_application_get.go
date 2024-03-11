package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apm_application_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a New Relic application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apm_application_getCmd).Standalone()
	apm_applicationCmd.AddCommand(apm_application_getCmd)
}
