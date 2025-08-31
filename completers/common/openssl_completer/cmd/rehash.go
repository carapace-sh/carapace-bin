package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var rehashCmd = &cobra.Command{
	Use:     "rehash",
	Short:   "Create symbolic links to certificate and CRL files named by the hash values",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rehashCmd).Standalone()

	rehashCmd.Flags().BoolS("compat", "compat", false, "Create both new- and old-style hash links")
	rehashCmd.Flags().BoolS("h", "h", false, "Display this summary")
	rehashCmd.Flags().BoolS("n", "n", false, "Do not remove existing links")
	rehashCmd.Flags().BoolS("old", "old", false, "Use old-style hash to generate links")
	rehashCmd.Flags().BoolS("v", "v", false, "Verbose output")
	common.AddProviderFlags(rehashCmd)
	rootCmd.AddCommand(rehashCmd)

	carapace.Gen(rehashCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
