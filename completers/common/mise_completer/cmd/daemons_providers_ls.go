package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_providers_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "[experimental] List shared servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_providers_lsCmd).Standalone()

	daemons_providers_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	daemons_providers_lsCmd.Flags().Bool("json", false, "Output in JSON format")
	daemons_providersCmd.AddCommand(daemons_providers_lsCmd)
}
