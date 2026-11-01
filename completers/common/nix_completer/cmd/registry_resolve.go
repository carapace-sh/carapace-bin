package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var registry_resolveCmd = &cobra.Command{
	Use:   "resolve [flags] [flake-refs...]",
	Short: "resolve flake references using the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_resolveCmd).Standalone()

	common.AddLoggingFlags(registry_resolveCmd)

	registryCmd.AddCommand(registry_resolveCmd)

	carapace.Gen(registry_resolveCmd).PositionalAnyCompletion(nix.ActionFlakeRefs())
}
