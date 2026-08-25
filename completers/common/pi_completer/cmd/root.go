package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pi",
	Short: "AI coding agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Gen(rootCmd).Standalone()
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().String("provider", "", "Provider name (default: google)")
	rootCmd.Flags().String("model", "", "Model pattern or ID")
	rootCmd.Flags().String("api-key", "", "API key (defaults to env vars)")
	rootCmd.Flags().String("system-prompt", "", "System prompt")
	rootCmd.Flags().String("append-system-prompt", "", "Append text or file contents to the system prompt")

	rootCmd.Flags().String("mode", "", "Output mode")

	rootCmd.Flags().BoolP("print", "p", false, "Non-interactive mode: process prompt and exit")
	rootCmd.Flags().BoolP("continue", "c", false, "Continue previous session")
	rootCmd.Flags().BoolP("resume", "r", false, "Select a session to resume")

	rootCmd.Flags().String("session", "", "Use specific session file or partial UUID")
	rootCmd.Flags().String("session-id", "", "Use exact project session ID")
	rootCmd.Flags().String("fork", "", "Fork specific session into a new session")
	rootCmd.Flags().String("session-dir", "", "Directory for session storage and lookup")
	rootCmd.Flags().Bool("no-session", false, "Don't save session (ephemeral)")

	rootCmd.Flags().StringP("name", "n", "", "Set session display name")
	rootCmd.Flags().String("models", "", "Comma-separated model patterns for Ctrl+P cycling")

	rootCmd.Flags().StringP("tools", "t", "", "Comma-separated allowlist of tool names to enable")

	rootCmd.Flags().String("thinking", "", "Set thinking level")

	rootCmd.Flags().StringP("extension", "e", "", "Load an extension file")

	rootCmd.Flags().String("skill", "", "Load a skill file or directory")

	rootCmd.Flags().String("prompt-template", "", "Load a prompt template file or directory")

	rootCmd.Flags().String("theme", "", "Load a theme file or directory")
	rootCmd.Flags().String("use-theme", "", "Set the initial interactive theme")

	rootCmd.Flags().String("export", "", "Export session file to HTML and exit")

	rootCmd.Flags().Bool("list-models", false, "List available models (with optional fuzzy search)")
	rootCmd.Flags().Bool("verbose", false, "Force verbose startup")

	rootCmd.Flags().String("tui-mode", "", "TUI mode")

	rootCmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this run")

	rootCmd.Flags().Bool("offline", false, "Disable startup network operations")

	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().BoolP("version", "v", false, "Show version number")

	embed.SubcommandsAsFlags(rootCmd,
		&cobra.Command{
			Use:     "no-tools",
			Aliases: []string{"nt"},
			Short:   "Disable all tools by default",
		},
		&cobra.Command{
			Use:     "no-builtin-tools",
			Aliases: []string{"nbt"},
			Short:   "Disable built-in tools by default but keep extension/custom tools",
		},
		&cobra.Command{
			Use:     "exclude-tools",
			Aliases: []string{"xt"},
			Short:   "Comma-separated denylist of tool names to disable",
		},
		&cobra.Command{
			Use:     "no-extensions",
			Aliases: []string{"ne"},
			Short:   "Disable extension discovery",
		},
		&cobra.Command{
			Use:     "no-skills",
			Aliases: []string{"ns"},
			Short:   "Disable skills discovery and loading",
		},
		&cobra.Command{
			Use:     "no-prompt-templates",
			Aliases: []string{"np"},
			Short:   "Disable prompt template discovery",
		},
		&cobra.Command{
			Use:     "no-context-files",
			Aliases: []string{"nc"},
			Short:   "Disable AGENTS.md and CLAUDE.md discovery",
		},
		&cobra.Command{
			Use:     "no-approve",
			Aliases: []string{"na"},
			Short:   "Ignore project-local files for this run",
		},
		&cobra.Command{
			Use:   "no-themes",
			Short: "Disable theme discovery and loading",
		},
	)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValues(
			"text",
			"json",
			"rpc",
		),

		"session":     carapace.ActionFiles(),
		"fork":        carapace.ActionFiles(),
		"session-dir": carapace.ActionDirectories(),

		"thinking": carapace.ActionValues(
			"off",
			"minimal",
			"low",
			"medium",
			"high",
			"xhigh",
			"max",
		),

		"extension":       carapace.ActionFiles(),
		"skill":           carapace.ActionFiles(),
		"prompt-template": carapace.ActionFiles(),
		"theme":           carapace.ActionFiles(),
		"export":          carapace.ActionFiles(),

		"tui-mode": carapace.ActionValues(
			"regular",
			"fullscreen",
		),
	})

	rootCmd.AddCommand(
		installCmd(),
		removeCmd(),
		updateCmd(),
		listCmd(),
		configCmd(),
		authCmd(),
	)
}

func installCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install extension source and add to settings",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().BoolP("local", "l", false, "Install project-locally (.pi/settings.json)")
	commandFlags(cmd)

	return cmd
}

func removeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"uninstall"},
		Short:   "Remove extension source from settings",
		Run:    func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().BoolP("local", "l", false, "Remove from project settings (.pi/settings.json)")
	commandFlags(cmd)

	return cmd
}

func updateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update pi, extensions, or model catalogs",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().Bool("self", false, "Update pi only")
	cmd.Flags().Bool("extensions", false, "Update installed packages only")
	cmd.Flags().Bool("models", false, "Refresh model catalogs only")
	cmd.Flags().Bool("all", false, "Update pi and installed packages")
	cmd.Flags().String("extension", "", "Update one package only")
	cmd.Flags().Bool("force", false, "Reinstall pi even if the current version is latest")

	commandFlags(cmd)

	return cmd
}

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List installed extensions from settings",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	commandFlags(cmd)

	return cmd
}

func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Open TUI to enable/disable package resources",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().BoolP("local", "l", false, "Edit project overrides (.pi/settings.json)")
	commandFlags(cmd)

	return cmd
}

func commandFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this command")

	embed.SubcommandsAsFlags(cmd,
		&cobra.Command{
			Use:     "no-approve",
			Aliases: []string{"na"},
			Short:   "Ignore project-local files for this command",
		},
	)
}

func authCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Print credentials for external clients",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.AddCommand(
		authCheckCmd(),
		authPrintAPIKeyCmd(),
		authPrintBearerTokenCmd(),
	)

	return cmd
}

func authCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check provider readiness",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().String("provider", "", "Provider name")
	cmd.Flags().String("model", "", "Model pattern or ID")
	cmd.Flags().Bool("json", false, "Output result as JSON")
	cmd.Flags().Bool("credentials", false, "Emit the credential (or include it in JSON)")
	cmd.Flags().Bool("no-refresh", false, "Do not refresh expired OAuth credentials")

	return cmd
}

func authPrintAPIKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print-api-key",
		Short: "Print a provider API key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().String("provider", "", "Provider name")
	cmd.Flags().String("model", "", "Model pattern or ID")

	return cmd
}

func authPrintBearerTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print-bearer-token",
		Short: "Print an OAuth bearer token",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().String("provider", "", "Provider name")
	cmd.Flags().String("model", "", "Model pattern or ID")
	cmd.Flags().String("min-expiry", "", "Minimum expiry duration (e.g. 30m, 1h)")

	return cmd
}