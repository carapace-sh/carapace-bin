// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Bootstrap one or more machines over SSH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_remoteCmd).Standalone()

	bootstrap_remoteCmd.Flags().String("adopt", "", "Adopt global configuration or shared dotfile history on each target")
	bootstrap_remoteCmd.Flags().Bool("all", false, "Select every inventory host")
	bootstrap_remoteCmd.Flags().String("bootstrap-command", "", "Shell command, run on each target, that installs mise and puts it on PATH")
	bootstrap_remoteCmd.Flags().String("connect-timeout", "10", "SSH connection timeout in seconds")
	bootstrap_remoteCmd.Flags().StringSlice("copy-link", nil, "Copy the target of this source-relative symbolic link instead of the link; repeat for more")
	bootstrap_remoteCmd.Flags().Bool("copy-links", false, "Copy the targets of every symbolic link in the source instead of the links")
	bootstrap_remoteCmd.Flags().BoolP("dry-run", "n", false, "Show what would change on each target without changing it")
	bootstrap_remoteCmd.Flags().StringSlice("exclude", nil, "Another pattern to leave out of the archive; repeat for more")
	bootstrap_remoteCmd.Flags().Bool("fail-fast", false, "Stop after the first failed target")
	bootstrap_remoteCmd.Flags().Bool("force-dotfiles", false, "Overwrite files on each target that conflict with whole-file dotfile entries")
	bootstrap_remoteCmd.Flags().Bool("github-relay-all-repos", false, "Let targets read every repository your local credentials can read (needs `--github-relay-read-only`)")
	bootstrap_remoteCmd.Flags().String("github-relay-log-format", "", "Relay log and summary format, `text` or `jsonl` (needs `--github-relay-read-only`)")
	bootstrap_remoteCmd.Flags().Bool("github-relay-log-requests", false, "Log sanitized relay requests on local stderr (needs `--github-relay-read-only`)")
	bootstrap_remoteCmd.Flags().String("github-relay-max-duration", "", "End borrowed access after a duration such as `1h`; `0s` lasts the whole run (needs `--github-relay-read-only`)")
	bootstrap_remoteCmd.Flags().Bool("github-relay-no-log-requests", false, "Turn off request logging, overriding the `github_relay` settings (needs `--github-relay-read-only`)")
	bootstrap_remoteCmd.Flags().Bool("github-relay-read-only", false, "Borrow read-only GitHub access for this run (Linux and macOS only)")
	bootstrap_remoteCmd.Flags().StringSlice("github-relay-repo", nil, "Repository the targets may read through the relay; repeat for more (needs `--github-relay-read-only`)")
	bootstrap_remoteCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_remoteCmd.Flags().StringSlice("host", nil, "SSH destination that is not in the inventory (`[user@]host`); repeat for more")
	bootstrap_remoteCmd.Flags().StringP("identity-file", "i", "", "SSH identity file, overriding the inventory and SSH config")
	bootstrap_remoteCmd.Flags().String("install-mise", "", "Install the provisioned mise on each host instead of only staging it")
	bootstrap_remoteCmd.Flags().Bool("keep-staging", false, "Keep the remote staging directory for debugging")
	bootstrap_remoteCmd.Flags().String("mise-bin", "", "Upload this local mise executable, for platforms without an official release binary")
	bootstrap_remoteCmd.Flags().Bool("no-install-mise", false, "Do not install mise on the host, even when the selected hosts configure it")
	bootstrap_remoteCmd.Flags().StringSlice("only", nil, "Run only these parts on each target")
	bootstrap_remoteCmd.Flags().String("port", "", "SSH port, overriding the inventory and SSH config")
	bootstrap_remoteCmd.Flags().Bool("prompt-secrets", false, "Prompt securely for missing secret inputs on the remote host")
	bootstrap_remoteCmd.Flags().StringSlice("remote-env", nil, "Config environments to load on each target; repeat or separate with commas (for example, ci,dotfiles)")
	bootstrap_remoteCmd.Flags().String("remote-mise", "", "Existing mise executable name or path; relative paths use the staged project")
	bootstrap_remoteCmd.Flags().StringSlice("skip", nil, "Skip these parts on each target")
	bootstrap_remoteCmd.Flags().String("source", "", "Local directory to archive and send to each target")
	bootstrap_remoteCmd.Flags().StringSlice("ssh-option", nil, "OpenSSH `-o` option; repeat for multiple options")
	bootstrap_remoteCmd.Flags().StringSlice("tag", nil, "Select configured hosts with this tag; repeat to match any tag")
	bootstrap_remoteCmd.Flags().Bool("update", false, "Refresh package metadata and update repositories on each target")
	bootstrap_remoteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts on each target")
	bootstrap_remoteCmd.Flag("install-mise").NoOptDefVal = " "
	bootstrapCmd.AddCommand(bootstrap_remoteCmd)

	carapace.Gen(bootstrap_remoteCmd).FlagCompletion(carapace.ActionMap{
		"copy-link":     carapace.ActionFiles(),
		"identity-file": carapace.ActionFiles(),
		"install-mise":  carapace.ActionFiles(),
		"mise-bin":      carapace.ActionFiles(),
		"only":          carapace.ActionValues("plugins", "packages", "accounts", "files", "services", "firewall", "compose", "repos", "dotfiles", "mise-shell-activate", "macos-defaults", "macos-launchd-agents", "linux-systemd-units", "user", "tools", "task", "final-hook", "shell", "defaults", "launchd", "systemd"),
		"skip":          carapace.ActionValues("plugins", "packages", "accounts", "files", "services", "firewall", "compose", "repos", "dotfiles", "mise-shell-activate", "macos-defaults", "macos-launchd-agents", "linux-systemd-units", "user", "tools", "task", "final-hook", "shell", "defaults", "launchd", "systemd"),
		"source":        carapace.ActionDirectories(),
	})
}
