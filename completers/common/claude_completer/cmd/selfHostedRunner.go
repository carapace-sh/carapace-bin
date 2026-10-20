package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfHostedRunnerCmd = &cobra.Command{
	Use:   "self-hosted-runner",
	Short: "Run a self-hosted environment runner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfHostedRunnerCmd).Standalone()

	selfHostedRunnerCmd.Flags().String("api-url", "", "API base URL")
	selfHostedRunnerCmd.Flags().String("base-dir", "", "Base directory for repo checkouts (default: /workspace;")
	selfHostedRunnerCmd.Flags().String("capacity", "", "Max concurrent sessions")
	selfHostedRunnerCmd.Flags().String("client-label", "", "Observability label sent at registration")
	selfHostedRunnerCmd.Flags().Bool("configure-git", false, "Set global git identity to Claude <noreply@anthropic.com> and enable")
	selfHostedRunnerCmd.Flags().String("confine-repo-settings", "", "confine-repo-settings")
	selfHostedRunnerCmd.Flags().String("defer-shutdown-max-min", "", "defer-shutdown-max-min")
	selfHostedRunnerCmd.Flags().String("drain-grace-sec", "", "Default: 0 — exit immediately after active sessions finish, WITHOUT")
	selfHostedRunnerCmd.Flags().String("drain-marker-file", "", "For hosts that announce a graceful drain by writing a marker file")
	selfHostedRunnerCmd.Flags().String("drain-wait-sec", "", "On SIGTERM/SIGINT, wait up to N seconds for each session's in-flight")
	selfHostedRunnerCmd.Flags().String("environment-secret-file", "", "environment-secret-file")
	selfHostedRunnerCmd.Flags().String("exec-path", "", "Binary to spawn for child sessions")
	selfHostedRunnerCmd.Flags().String("exit-if-unused-min", "", "Exit the runner if never assigned work for N min (autoscaler scale-down)")
	selfHostedRunnerCmd.Flags().Bool("git-host-rewrite", false, "<f>=<t> Rewrite https://<f>/")
	selfHostedRunnerCmd.Flags().String("git-ssh-rewrite", "", "Rewrite https://<host>/")
	selfHostedRunnerCmd.Flags().String("health-port", "", "Port for /healthz HTTP listener")
	selfHostedRunnerCmd.Flags().BoolP("help", "h", false, "Display help for command")
	selfHostedRunnerCmd.Flags().Bool("help,", false, "-h Show this help message")
	selfHostedRunnerCmd.Flags().String("hooks-dir", "", "Directory of lifecycle hook scripts (checkout, command, post-session)")
	selfHostedRunnerCmd.Flags().String("host-config-snapshot", "", "host-config-snapshot")
	selfHostedRunnerCmd.Flags().String("kill-session-after-min", "", "Cap a session child's wall-clock life at N min (runaway backstop)")
	selfHostedRunnerCmd.Flags().String("lock-to-account", "", "Lock runner to a single account at registration (webhook-driven on-demand")
	selfHostedRunnerCmd.Flags().String("log-file", "", "Tee runner logs to a file in append mode")
	selfHostedRunnerCmd.Flags().String("log-level", "", "Log level: info or debug")
	selfHostedRunnerCmd.Flags().String("post-session-hook-timeout-sec", "", "post-session-hook-timeout-sec")
	selfHostedRunnerCmd.Flags().String("proxy-authorization-command", "", "proxy-authorization-command")
	selfHostedRunnerCmd.Flags().String("proxy-authorization-file", "", "proxy-authorization-file")
	selfHostedRunnerCmd.Flags().Bool("push-outcome-on-release", false, "On a runner-initiated non-completed session end (SIGTERM drain,")
	selfHostedRunnerCmd.Flags().String("release-idle-session-min", "", "Release a session slot after N min of no user input (turn finished,")
	selfHostedRunnerCmd.Flags().String("remove-session-state", "", "remove-session-state")
	selfHostedRunnerCmd.Flags().String("retire-at", "", "Retire the runner at the given wall-clock time (absolute Unix timestamp, in")
	selfHostedRunnerCmd.Flags().String("server-auto-mode-lists", "", "server-auto-mode-lists")
	selfHostedRunnerCmd.Flags().String("session-stop-grace-sec", "", "session-stop-grace-sec")
	selfHostedRunnerCmd.Flags().String("startup-timeout-min", "", "Release a session slot if the child has not completed initialization")
	selfHostedRunnerCmd.Flags().String("trust-workspace", "", "Seed persisted trust for each session's repo paths so repo-level")
	selfHostedRunnerCmd.Flags().Bool("use-anthropic-git-proxy", false, "Opt this runner into Anthropic-managed git (clones are authenticated")
	selfHostedRunnerCmd.Flag("remove-session-state").NoOptDefVal = " "
	selfHostedRunnerCmd.Flag("trust-workspace").NoOptDefVal = " "
	rootCmd.AddCommand(selfHostedRunnerCmd)

	carapace.Gen(selfHostedRunnerCmd).FlagCompletion(carapace.ActionMap{
		"base-dir":                 carapace.ActionDirectories(),
		"drain-marker-file":        carapace.ActionFiles(),
		"environment-secret-file":  carapace.ActionFiles(),
		"exec-path":                carapace.ActionFiles(),
		"hooks-dir":                carapace.ActionDirectories(),
		"host-config-snapshot":     carapace.ActionValues("disk", "memory"),
		"log-file":                 carapace.ActionFiles(),
		"log-level":                carapace.ActionValues("info", "debug"),
		"proxy-authorization-file": carapace.ActionFiles(),
	})
}
