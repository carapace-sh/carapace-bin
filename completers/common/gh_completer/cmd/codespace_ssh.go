package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_sshCmd = &cobra.Command{
	Use:   "ssh [<flags>...] [-- <ssh-flags>...] [<command>]",
	Short: "SSH into a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_sshCmd).Standalone()

	codespace_sshCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_sshCmd.Flags().Bool("config", false, "Write OpenSSH configuration to stdout")
	codespace_sshCmd.Flags().BoolP("debug", "d", false, "Log debug data to a file")
	codespace_sshCmd.Flags().String("debug-file", "", "Path of the file log to")
	codespace_sshCmd.Flags().String("profile", "", "Name of the SSH profile to use")
	codespace_sshCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_sshCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespace_sshCmd.Flags().String("server-port", "", "SSH server port number (0 => pick unused)")
	codespace_sshCmd.Flags().Bool("stdio", false, "Proxy sshd connection to stdio")
	codespace_sshCmd.Flag("stdio").Hidden = true
	codespaceCmd.AddCommand(codespace_sshCmd)

	carapace.Gen(codespace_sshCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"debug-file": carapace.ActionFiles(),
		"profile":    carapace.ActionValues("TODO"), // TODO complete ssh profiles
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
	})

	// TODO ssh flags when os.Args contains `--`
}
