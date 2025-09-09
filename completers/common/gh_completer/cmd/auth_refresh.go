package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var auth_refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh stored authentication credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_refreshCmd).Standalone()

	auth_refreshCmd.Flags().BoolP("clipboard", "c", false, "Copy one-time OAuth device code to clipboard")
	auth_refreshCmd.Flags().StringP("hostname", "h", "", "The GitHub host to use for authentication")
	auth_refreshCmd.Flags().Bool("insecure-storage", false, "Save authentication credentials in plain text instead of credential store")
	auth_refreshCmd.Flags().StringSliceP("remove-scopes", "r", nil, "Authentication scopes to remove from gh")
	auth_refreshCmd.Flags().Bool("reset-scopes", false, "Reset authentication scopes to the default minimum set of scopes")
	auth_refreshCmd.Flags().StringSliceP("scopes", "s", nil, "Additional authentication scopes for gh to have")
	auth_refreshCmd.Flags().Bool("secure-storage", false, "Save authentication credentials in secure credential store")
	auth_refreshCmd.Flag("secure-storage").Hidden = true
	authCmd.AddCommand(auth_refreshCmd)

	carapace.Gen(auth_refreshCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"remove-scopes": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			hostname := auth_refreshCmd.Flag("hostname").Value.String()
			return gh.ActionCurrentAuthScopes(hostname).MultiParts(":").UniqueList(",")
		}),
		"scopes": gh.ActionAuthScopes().MultiParts(":").UniqueList(","),
	})
}
