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
		"act": {
			Description: "Run GitHub actions locally",
			Url:         "https://github.com/nektos/act",
			Variant:     "cobra",
		},
		"ansible": {
			Description: "Define and run a single task ‘playbook’ against a set of hosts",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible.html",
			Variant:     "argcomplete",
		},
		"ansible-config": {
			Description: "View ansible configuration",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-config.html",
			Variant:     "argcomplete",
		},
		"ansible-console": {
			Description: "REPL console for executing Ansible tasks",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-console.html",
			Variant:     "argcomplete",
		},
		"ansible-creator": {
			Description: "The fastest way to generate all your ansible content",
			Url:         "https://docs.ansible.com/projects/creator/",
			Variant:     "argcomplete",
		},
		"ansible-doc": {
			Description: "Ansible plugin documentation tool",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-doc.html",
			Variant:     "argcomplete",
		},
		"ansible-galaxy": {
			Description: "Perform various Role and Collection related operations",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-galaxy.html",
			Variant:     "argcomplete",
		},
		"ansible-inventory": {
			Description: "Show Ansible inventory information",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-inventory.html",
			Variant:     "argcomplete",
		},
		"ansible-playbook": {
			Description: "Runs Ansible playbooks, executing the defined tasks on the targeted hosts",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-playbook.html",
			Variant:     "argcomplete",
		},
		"ansible-pull": {
			Description: "Pulls playbooks from a VCS repo and executes them on target host",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-pull.html",
			Variant:     "argcomplete",
		},
		"ansible-vault": {
			Description: "encryption/decryption utility for Ansible data files",
			Url:         "https://docs.ansible.com/projects/ansible/latest/cli/ansible-vault.html",
			Variant:     "argcomplete",
		},
		"apko": {
			Description: "Build OCI images from APK packages directly without Dockerfile",
			Url:         "https://github.com/chainguard-dev/apko",
			Variant:     "cobra",
		},
		"apptainer": {
			Description: "Application container and unprivileged sandbox platform for Linux",
			Url:         "https://apptainer.org/",
			Variant:     "cobra",
		},
		"aqua": {
			Description: "Declarative CLI Version manager",
			Url:         "https://aquaproj.github.io/",
			Variant:     "urfavecli",
		},
		"argo": {
			Description: "Workflow Engine for Kubernetes",
			Url:         "https://argoproj.github.io/",
			Variant:     "cobra",
		},
		"argocd": {
			Description: "Declarative Continuous Deployment for Kubernetes",
			Url:         "https://argoproj.github.io/cd/",
			Variant:     "cobra",
		},
		"az": {
			Description: "Azure Command-Line Interface",
			Url:         "https://docs.microsoft.com/en-us/cli/azure/",
			Variant:     "argcomplete",
		},
		"boundary": {
			Description: "Boundary enables identity-based access management for dynamic infrastructure",
			Url:         "https://www.boundaryproject.io/downloads",
			Variant:     "complete",
		},
		"cdebug": {
			Description: "A Swiss army knife of container debugging",
			Url:         "https://github.com/iximiuz/cdebug",
			Variant:     "cobra",
		},
		"cekit": {
			Description: "Container Evolution Kit",
			Url:         "https://cekit.io/",
			Variant:     "click",
		},
		"chezmoi": {
			Description: "Manage your dotfiles across multiple diverse machines, securely",
			Url:         "https://chezmoi.io/",
			Variant:     "cobra",
		},
		"colima": {
			Description: "Container runtimes on macOS (and Linux) with minimal setup",
			Url:         "https://colima.run/",
			Variant:     "cobra",
		},
		"cosign": {
			Description: "Code signing and transparency for containers and binaries",
			Url:         "https://github.com/sigstore/cosign",
			Variant:     "cobra",
		},
		"crc": {
			Description: "Local OpenShift 4.x cluster",
			Url:         "https://crc.dev/",
			Variant:     "cobra",
		},
		"cue": {
			Description: "CUE makes it easy to validate data",
			Url:         "https://github.com/cue-lang/cue",
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
		"doctl": {
			Description: "doctl is a command line interface (CLI) for the DigitalOcean API",
			Url:         "https://docs.digitalocean.com/reference/doctl/",
			Variant:     "cobra",
		},
		"exercism": {
			Description: "A command-line interface for Exercism",
			Url:         "https://exercism.org/",
			Variant:     "cobra",
		},
		"flyctl": {
			Description: "Command line tools for fly.io services",
			Url:         "https://fly.io/",
			Variant:     "cobra",
		},
		"ghalint": {
			Description: "GitHub Actions linter",
			Url:         "https://github.com/suzuki-shunsuke/ghalint",
			Variant:     "urfavecli",
		},
		"gitleaks": {
			Description: "Gitleaks scans code, past or present, for secrets",
			Url:         "https://gitleaks.io/",
			Variant:     "cobra",
		},
		"gitlint": {
			Description: "Linting for your git commit messages",
			Url:         "https://jorisroovers.com/gitlint/latest/",
			Variant:     "click",
		},
		"gitsign": {
			Description: "Keyless Git signing using Sigstore",
			Url:         "https://github.com/sigstore/gitsign",
			Variant:     "cobra",
		},
		"gomplate": {
			Description: "Process text files with Go templates",
			Url:         "https://gomplate.ca/",
			Variant:     "cobra",
		},
		"hatch": {
			Description: "Modern, extensible Python project management",
			Url:         "https://hatch.pypa.io/latest/",
			Variant:     "click",
		},
		"hcloud": {
			Description: "A CLI for Hetzner Cloud",
			Url:         "https://github.com/hetznercloud/cli",
			Variant:     "cobra",
		},
		"helmfile": {
			Description: "Declaratively deploy Kubernetes Helm Charts",
			Url:         "https://github.com/helmfile/helmfile",
			Variant:     "cobra",
		},
		"incus": {
			Description: "Powerful system container and virtual machine manager",
			Url:         "https://linuxcontainers.org/incus",
			Variant:     "cobra",
		},
		"iredis": {
			Description: "A Terminal Client for Redis with AutoCompletion and Syntax Highlighting",
			Url:         "https://iredis.xbin.io/",
			Variant:     "click",
		},
		"k3sup": {
			Description: "Bootstrap K3s over SSH in < 60s",
			Url:         "https://github.com/alexellis/k3sup",
			Variant:     "cobra",
		},
		"k6": {
			Description: "A modern load testing tool, using Go and JavaScript",
			Url:         "https://k6.io",
			Variant:     "cobra",
		},
		"k9s": {
			Description: "K9s is a CLI to view and manage your Kubernetes clusters",
			Url:         "https://k9scli.io/",
			Variant:     "cobra",
		},
		"kcl": {
			Description: "The KCL Command Line Interface",
			Url:         "https://github.com/kcl-lang/cli",
			Variant:     "cobra",
		},
		"kitten": {
			Description: "Fast, statically compiled implementations of various kittens",
			Url:         "https://sw.kovidgoyal.net/kitty/",
			Variant:     "kitten",
		},
		"kitty": {
			Description: "The fast, feature rich terminal emulator",
			Url:         "https://sw.kovidgoyal.net/kitty/",
			Variant:     "kitten",
		},
		"kubebuilder": {
			Description: "CLI tool for building Kubernetes extensions and tools",
			Url:         "https://github.com/kubernetes-sigs/kubebuilder",
			Variant:     "cobra",
		},
		"kustomize": {
			Description: "Manages declarative configuration of Kubernetes",
			Url:         "https://github.com/kubernetes-sigs/kustomize",
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
		"litecli": {
			Description: "CLI for SQLite Databases with auto-completion and syntax highlighting",
			Url:         "https://litecli.com/",
			Variant:     "click",
		},
		"mycli": {
			Description: "A MySQL terminal client with auto-completion and syntax highlighting",
			Url:         "https://www.mycli.net/",
			Variant:     "click",
		},
		"nerdctl": {
			Description: "Nerdctl is a command line interface for containerd",
			Url:         "https://github.com/containerd/nerdctl",
			Variant:     "cobra",
		},
		"nomad": {
			Description: "Nomad is an easy-to-use, flexible, and performant workload orchestrator",
			Url:         "https://www.nomadproject.io/",
			Variant:     "complete",
		},
		"nox": {
			Description: "Flexible test automation for Python",
			Url:         "https://nox.thea.codes/",
			Variant:     "argcomplete",
		},
		"oh-my-posh": {
			Description: "A cross platform tool to render your prompt",
			Url:         "https://ohmyposh.dev/",
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
		"pgcli": {
			Description: "Postgres CLI with autocompletion and syntax highlighting",
			Url:         "https://pgcli.com/",
			Variant:     "click",
		},
		"pipenv": {
			Description: "Python Development Workflow for Humans",
			Url:         "https://pipenv.pypa.io/",
			Variant:     "click",
		},
		"pipx": {
			Description: "Install and execute apps from Python packages",
			Url:         "https://pipx.pypa.io/",
			Variant:     "argcomplete",
		},
		"podman": {
			Description: "Simple management tool for pods, containers and images",
			Url:         "https://podman.io/",
			Variant:     "cobra",
		},
		"pscale": {
			Description: "A CLI for communicating with PlanetScale's API",
			Url:         "https://planetscale.com/docs/cli",
			Variant:     "cobra",
		},
		"ramalama": {
			Description: "tool for working with LLM models",
			Url:         "https://ramalama.ai/",
			Variant:     "argcomplete",
		},
		"rclone": {
			Description: "Rsync for cloud storage",
			Url:         "https://rclone.org/",
			Variant:     "cobra",
		},
		"reuse": {
			Description: "reuse is a tool for compliance with the REUSE recommendations",
			Url:         "https://reuse.software",
			Variant:     "click",
		},
		"slsa-verifier": {
			Description: "Verify provenance from SLSA compliant builders",
			Url:         "https://github.com/slsa-framework/slsa-verifier",
			Variant:     "cobra",
		},
		"talosctl": {
			Description: "A CLI for out-of-band management of Kubernetes nodes created by Talos",
			Url:         "https://docs.siderolabs.com/talos/latest/reference/cli",
			Variant:     "cobra",
		},
		"trivy": {
			Description: "Scanner for vulnerabilities in container images, file systems, and Git repos",
			Url:         "https://trivy.dev/",
			Variant:     "cobra",
		},
		"toolbox": {
			Description: "Tool for interactive command line environments on Linux",
			Url:         "https://containertoolbx.org/",
			Variant:     "cobra",
		},
		"vault": {
			Description: "A tool for secrets management",
			Url:         "https://www.vaultproject.io/",
			Variant:     "complete",
		},
		"vunnel": {
			Description: "Tool for collecting vulnerability data from various sources",
			Url:         "https://github.com/anchore/vunnel",
			Variant:     "click",
		},
		"ykman": {
			Description: "Configure your YubiKey via the command line",
			Url:         "https://developers.yubico.com/yubikey-manager/",
			Variant:     "click",
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
