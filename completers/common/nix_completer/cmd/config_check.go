package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var config_checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "check your system for potential problems and print a PASS or FAIL for each check",
	Aliases: []string{"doctor"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_checkCmd).Standalone()

	common.AddLoggingFlags(config_checkCmd)

	configCmd.AddCommand(config_checkCmd)
}
