package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var derivation_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a store derivation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(derivation_addCmd).Standalone()

	derivation_addCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")

	derivationCmd.AddCommand(derivation_addCmd)

	common.AddLoggingFlags(derivation_addCmd)
}
