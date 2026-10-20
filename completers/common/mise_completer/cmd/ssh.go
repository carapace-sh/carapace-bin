package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Open an SSH session, optionally borrowing read-only GitHub access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCmd).Standalone()

	sshCmd.Flags().Bool("github-relay-all-repos", false, "Allow reads of every repository your local GitHub credentials can access")
	sshCmd.Flags().String("github-relay-log-format", "", "Relay log and summary format: text or jsonl (default: the github_relay.log_format setting)")
	sshCmd.Flags().Bool("github-relay-log-requests", false, "Log sanitized relay requests on local stderr")
	sshCmd.Flags().String("github-relay-max-duration", "", "End borrowed access after a duration such as 1h; 0s lasts the whole session")
	sshCmd.Flags().Bool("github-relay-no-log-requests", false, "Turn off request logging even if the github_relay.log_requests setting is on")
	sshCmd.Flags().Bool("github-relay-read-only", false, "Borrow read-only GitHub access for this session only")
	sshCmd.Flags().StringSlice("github-relay-repo", nil, "GitHub repository the host may read; repeat for more repositories")
	sshCmd.Flags().Bool("global-config-directory", false, "")
	sshCmd.Flags().BoolP("help", "h", false, "Print help")
	sshCmd.Flags().StringP("identity-file", "i", "", "SSH identity file")
	sshCmd.Flags().StringP("port", "p", "", "SSH port")
	sshCmd.Flags().String("relay-session", "", "Internal session adapter")
	sshCmd.Flags().String("repository-bundle", "", "")
	sshCmd.Flags().Bool("repository-dry-run", false, "Preview the transferred repository's installation without writing it")
	sshCmd.Flags().String("repository-origin", "", "")
	sshCmd.Flags().String("repository-preview-directory", "", "")
	sshCmd.Flags().Bool("repository-require-branch", false, "The bundled branch was requested with `?ref=`")
	sshCmd.Flags().String("repository-revision", "", "")
	sshCmd.Flags().Bool("repository-update", false, "")
	sshCmd.Flags().Bool("repository-yes", false, "")
	sshCmd.Flags().StringSliceP("ssh-option", "o", nil, "OpenSSH option; repeat for multiple options")
	sshCmd.Flag("global-config-directory").Hidden = true
	sshCmd.Flag("relay-session").Hidden = true
	sshCmd.Flag("repository-bundle").Hidden = true
	sshCmd.Flag("repository-dry-run").Hidden = true
	sshCmd.Flag("repository-origin").Hidden = true
	sshCmd.Flag("repository-preview-directory").Hidden = true
	sshCmd.Flag("repository-require-branch").Hidden = true
	sshCmd.Flag("repository-revision").Hidden = true
	sshCmd.Flag("repository-update").Hidden = true
	sshCmd.Flag("repository-yes").Hidden = true
	rootCmd.AddCommand(sshCmd)
}
