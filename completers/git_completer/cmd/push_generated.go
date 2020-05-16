package cmd

import (
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Update remote refs along with associated objects",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	pushCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	pushCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	pushCmd.Flags().Bool("all", false, "push all refs")
	pushCmd.Flags().Bool("atomic", false, "request atomic transaction on remote side")
	pushCmd.Flags().BoolP("delete", "d", false, "delete refs")
	pushCmd.Flags().String("exec", "", "receive pack program")
	pushCmd.Flags().BoolP("force", "f", false, "force updates")
	pushCmd.Flags().Bool("follow-tags", false, "push missing but relevant tags")
	pushCmd.Flags().String("force-with-lease", "", "require old value of ref to be at this value")
	pushCmd.Flags().Bool("mirror", false, "mirror all refs")
	pushCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	pushCmd.Flags().Bool("no-verify", false, "bypass pre-push hook")
	pushCmd.Flags().BoolP("push-option", "o", false, "<server-specific>    option to transmit")
	pushCmd.Flags().Bool("porcelain", false, "machine-readable output")
	pushCmd.Flags().Bool("progress", false, "force progress reporting")
	pushCmd.Flags().Bool("prune", false, "prune locally removed refs")
	pushCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	pushCmd.Flags().String("receive-pack", "", "receive pack program")
	pushCmd.Flags().String("recurse-submodules", "", "control recursive pushing of submodules")
	pushCmd.Flags().String("repo", "", "repository")
	pushCmd.Flags().String("signed", "", "GPG sign the push")
	pushCmd.Flags().Bool("tags", false, "push tags (can't be used with --all or --mirror)")
	pushCmd.Flags().Bool("thin", false, "use thin pack")
	pushCmd.Flags().BoolP("set-upstream", "u", false, "set upstream for git pull/status")
	pushCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	rootCmd.AddCommand(pushCmd)
}
