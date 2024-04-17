package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_makeContentAddressedCmd = &cobra.Command{
	Use:   "make-content-addressed",
	Short: "rewrite a path or closure to content-addressed form",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_makeContentAddressedCmd).Standalone()

	store_makeContentAddressedCmd.Flags().String("from", "", "URL of the source Nix store")
	store_makeContentAddressedCmd.Flags().Bool("json", false, "Produce output in JSON format")
	store_makeContentAddressedCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	store_makeContentAddressedCmd.Flags().String("to", "", "URL of the destination Nix store")
	storeCmd.AddCommand(store_makeContentAddressedCmd)

	addEvaluationFlags(store_makeContentAddressedCmd)
	addFlakeFlags(store_makeContentAddressedCmd)
	addLoggingFlags(store_makeContentAddressedCmd)

	// TODO flag completion
	carapace.Gen(store_makeContentAddressedCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
