package completers

import (
	"errors"

	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

func knownbridges() map[string]*completer.Completer {
	// TODO what about different variants of a command
	// TODO limit to target os
	m := map[string]*completer.Completer{
		"apptainer": {
			Description: "Application container and unprivileged sandbox platform for Linux",
			Url:         "https://apptainer.org/",
			Variant:     "cobra",
		},
		"chezmoi": {
			Description: "Manage your dotfiles across multiple diverse machines, securely",
			Url:         "https://chezmoi.io/",
			Variant:     "cobra",
		},
		"crc": {
			Description: "Local OpenShift 4.x cluster",
			Url:         "https://crc.dev/",
			Variant:     "cobra",
		},
		"devcontainer": {
			Description: "Use a Docker container as a development environment",
			Url:         "https://containers.dev/",
			Variant:     "yargs",
		},
		"devpod": {
			Description: "Codespaces but open-source, client-only and unopinionated",
			Url:         "https://devpod.sh/",
			Variant:     "cobra",
		},
		"flyctl": {
			Description: "Command line tools for fly.io services",
			Url:         "https://fly.io/",
			Variant:     "cobra",
		},
		"incus": {
			Description: "Powerful system container and virtual machine manager",
			Url:         "https://linuxcontainers.org/incus",
			Variant:     "cobra",
		},
		"lefthook": {
			Description: "Git hooks manager",
			Url:         "https://lefthook.dev/",
			Variant:     "urfavecli",
		},
		"limactl": {
			Description: "Linux virtual machines",
			Url:         "https://lima-vm.io/",
			Variant:     "cobra",
		},
		"nerdctl": {
			Description: "Nerdctl is a command line interface for containerd",
			Url:         "https://github.com/containerd/nerdctl",
			Variant:     "cobra",
		},
		"op": {
			Description: "1Password CLI",
			Url:         "https://developer.1password.com/docs/cli",
			Variant:     "cobra",
		},
		"orbctl": {
			// TODO darwin
			Description: "Manage OrbStack and its machines",
			Url:         "https://orbstack.dev/",
			Variant:     "cobra",
		},
		"podman": {
			Description: "Simple management tool for pods, containers and images",
			Url:         "https://podman.io/",
			Variant:     "cobra",
		},
		"rclone": {
			Description: "Rsync for cloud storage",
			Url:         "https://rclone.org/",
			Variant:     "cobra",
		},
	}

	for name, c := range m {
		c.Name = name
		c.Group = "bridge"
		c.Execute = func() error {
			if b, ok := bridge.Get(c.Variant); ok {
				return complete(name, b(name))()
			}
			return errors.New("unknown bridge: " + c.Variant)
		}
	}
	return m
}
