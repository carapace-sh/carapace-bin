package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display active account and authentication state on each known GitHub host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_statusCmd).Standalone()

	auth_statusCmd.Flags().BoolP("active", "a", false, "Display the active account only")
	auth_statusCmd.Flags().StringP("hostname", "h", "", "Check only a specific hostname's auth status")
	auth_statusCmd.Flags().String("jq", "", "Filter JSON output using a jq `expression`")
	auth_statusCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	auth_statusCmd.Flags().BoolP("show-token", "t", false, "Display the auth token")
	auth_statusCmd.Flags().String("template", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	authCmd.AddCommand(auth_statusCmd)

	carapace.Gen(auth_statusCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"json":     carapace.ActionValues("hosts"),
	})
}
