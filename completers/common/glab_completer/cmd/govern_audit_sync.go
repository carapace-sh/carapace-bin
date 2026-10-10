package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var govern_audit_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync agent session data to GitLab. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(govern_audit_syncCmd).Standalone()

	govern_audit_syncCmd.Flags().Bool("all", false, "Sync all sessions recorded by the hooks. Used by the fallback periodic sync.")
	govern_audit_syncCmd.Flags().Bool("complete", false, "Mark the session as completed. Used by the SessionEnd hook.")
	govern_audit_syncCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	govern_audit_syncCmd.Flags().Bool("silent", false, "Suppress all output. Used when invoked from hooks.")
	govern_auditCmd.AddCommand(govern_audit_syncCmd)

	carapace.Gen(govern_audit_syncCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(govern_audit_syncCmd),
	})
}
