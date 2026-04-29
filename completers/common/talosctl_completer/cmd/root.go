package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "talosctl",
	Short: "A CLI for out-of-band management of Kubernetes nodes created by Talos",
	Long:  "https://docs.siderolabs.com/talos/latest/reference/cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("cluster", "c", "", "Cluster to connect to if a proxy endpoint is used")
	rootCmd.PersistentFlags().String("context", "", "Context to be used in command")
	rootCmd.PersistentFlags().StringSliceP("endpoints", "e", nil, "Override default endpoints in Talos configuration")
	rootCmd.PersistentFlags().StringSliceP("nodes", "n", nil, "Target the specified nodes")
	rootCmd.PersistentFlags().String("siderov1-keys-dir", "", "Path to the SideroV1 auth PGP keys directory")
	rootCmd.PersistentFlags().String("talosconfig", "", "Path to the Talos configuration file")
	rootCmd.PersistentFlags().String("timeout", "", "Timeout")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"context":           actionContexts(),
		"siderov1-keys-dir": carapace.ActionDirectories(),
		"talosconfig":       carapace.ActionFiles(),
	})

	addCommands(rootCmd,
		"apply-config", "bootstrap", "cgroups", "containers", "copy", "dashboard", "debug", "dmesg", "edit", "events", "get", "health", "kubeconfig", "list", "logs", "memory", "mounts", "netstat", "patch", "pcap", "processes", "read", "reboot", "reset", "restart", "rollback", "rotate-ca", "service", "shutdown", "stats", "support", "time", "upgrade", "upgrade-k8s", "usage", "validate", "version",
	)
	addCommands(rootCmd, "cluster create dev", "cluster create docker", "cluster create qemu", "cluster destroy", "cluster show")
	addCommands(rootCmd, "completion bash", "completion fish", "completion powershell", "completion zsh")
	addCommands(rootCmd, "config add", "config context", "config contexts", "config endpoint", "config info", "config merge", "config new", "config node", "config remove")
	addCommands(rootCmd, "conformance kubernetes")
	addCommands(rootCmd, "etcd alarm disarm", "etcd alarm list", "etcd defrag", "etcd downgrade cancel", "etcd downgrade enable", "etcd downgrade validate", "etcd forfeit-leadership", "etcd leave", "etcd members", "etcd remove-member", "etcd snapshot", "etcd status")
	addCommands(rootCmd, "gen ca", "gen config", "gen crt", "gen csr", "gen key", "gen keypair", "gen secrets", "gen secureboot database", "gen secureboot pcr", "gen secureboot uki")
	addCommands(rootCmd, "image cache-cert-gen", "image cache-create", "image cache-serve", "image k8s-bundle", "image list", "image pull", "image remove", "image talos-bundle")
	addCommands(rootCmd, "inject serviceaccount")
	addCommands(rootCmd, "inspect dependencies")
	addCommands(rootCmd, "machineconfig gen", "machineconfig patch")
	addCommands(rootCmd, "meta delete", "meta write")
	addCommands(rootCmd, "wipe disk")

	carapace.Gen(rootCmd).PositionalCompletion(actionSubcommands(rootCmd))
}

func addCommands(root *cobra.Command, paths ...string) {
	for _, path := range paths {
		parent := root
		parts := strings.Fields(path)
		for i, part := range parts {
			cmd := findOrCreate(parent, part)
			if i == len(parts)-1 {
				cmd.Short = path
			}
			parent = cmd
		}
	}
}

func findOrCreate(parent *cobra.Command, name string) *cobra.Command {
	for _, cmd := range parent.Commands() {
		if cmd.Name() == name {
			return cmd
		}
	}
	cmd := &cobra.Command{Use: name, Short: name, Run: func(cmd *cobra.Command, args []string) {}}
	parent.AddCommand(cmd)
	carapace.Gen(cmd).PositionalCompletion(actionSubcommands(cmd))
	addLocalFlags(cmd)
	return cmd
}

func addLocalFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("file", "f", "", "File path")
	cmd.Flags().StringP("mode", "m", "", "Mode")
	cmd.Flags().BoolP("insecure", "i", false, "Use the insecure maintenance service")
	cmd.Flags().Bool("dry-run", false, "Preview the operation")
	cmd.Flags().StringSliceP("config-patch", "p", nil, "Config patches to apply")
	cmd.Flags().String("format", "", "Output format")
	cmd.Flags().String("output", "", "Output path or format")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"format": carapace.ActionValues("json", "yaml", "table"),
		"mode":   carapace.ActionValues("auto", "no-reboot", "reboot", "staged", "try"),
		"output": carapace.ActionFiles(),
	})
}

func actionSubcommands(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		for _, sub := range cmd.Commands() {
			if !sub.Hidden {
				vals = append(vals, sub.Name(), sub.Short)
			}
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForKeyword)
	})
}

func actionContexts() carapace.Action {
	return carapace.ActionExecCommand("talosctl", "config", "contexts")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0, len(lines))
		for _, line := range lines {
			fields := strings.Fields(strings.TrimPrefix(line, "*"))
			if len(fields) > 0 && fields[0] != "CURRENT" && fields[0] != "NAME" {
				vals = append(vals, fields[0])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
