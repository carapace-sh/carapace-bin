package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddProviderFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("propquery", "propquery", "", "Property query used when fetching algorithms")
	cmd.Flags().StringSliceS("provider", "provider", nil, "Provider to load (can be specified multiple times)")
	cmd.Flags().StringS("provider-path", "provider-path", "", "Provider load path (must be before 'provider' argument if required)")
	cmd.Flags().StringS("provparam", "provparam", "", "Set a provider key-value parameter")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"provider":      carapace.ActionValues("base", "default", "legacy"),
		"provider-path": carapace.ActionDirectories(),
	})
}
