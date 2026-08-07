package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_credentialHelperCmd = &cobra.Command{
	Use:    "credential-helper [flags]",
	Short:  "Implements a generic credential helper.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_credentialHelperCmd).Standalone()

	auth_credentialHelperCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	authCmd.AddCommand(auth_credentialHelperCmd)

	carapace.Gen(auth_credentialHelperCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(auth_credentialHelperCmd),
	})
}
