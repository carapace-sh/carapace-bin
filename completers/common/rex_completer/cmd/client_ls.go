package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var clientLsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List clients",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clientLsCmd).Standalone()

	clientLsCmd.Flags().Bool("all", false, "include command line clients")
	clientLsCmd.Flags().Bool("json", false, "output clients as JSON")
	clientLsCmd.Flags().BoolP("quiet", "q", false, "print IDs only, one per line")
	clientLsCmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	clientLsCmd.Flags().Bool("short", true, "show short IDs (--short=false shows full IDs)")

	clientCmd.AddCommand(clientLsCmd)

	carapace.Gen(clientLsCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(clientLsCmd),
	})
}
